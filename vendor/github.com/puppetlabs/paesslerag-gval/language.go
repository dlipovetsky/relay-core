package gval

import (
	"context"
	"fmt"
	"text/scanner"
	"unicode"
)

// Language is an expression language
type Language struct {
	prefixes        map[interface{}]prefix
	operators       map[string]operator
	operatorSymbols map[rune]struct{}
	init            func(context.Context, *Parser) (Evaluable, error)
	alternate       func(context.Context, *Parser) (Evaluable, error)
	selector        func(Evaluables) Evaluable
}

// NewLanguage returns the union of given Languages as new Language.
func NewLanguage(bases ...Language) Language {
	l := newLanguage()
	for _, base := range bases {
		for i, e := range base.prefixes {
			l.prefixes[i] = e
		}
		for i, e := range base.operators {
			l.operators[i] = e.merge(l.operators[i])
			l.operators[i].initiate(i)
		}
		for i := range base.operatorSymbols {
			l.operatorSymbols[i] = struct{}{}
		}
		if base.init != nil {
			l.init = base.init
		}
		if base.alternate != nil {
			l.alternate = base.alternate
		}
		if base.selector != nil {
			l.selector = base.selector
		}
	}
	return l
}

func newLanguage() Language {
	return Language{
		prefixes:        map[interface{}]prefix{},
		operators:       map[string]operator{},
		operatorSymbols: map[rune]struct{}{},
	}
}

func (l Language) ChainEvaluable(c context.Context, p *Parser) (Evaluable, error) {
	p.pushLanguage(l)
	defer p.popLanguage()

	init := l.init
	if init == nil {
		init = func(c context.Context, p *Parser) (Evaluable, error) {
			return p.ParseExpression(c)
		}
	}

	return init(c, p)
}

// NewEvaluable returns an Evaluable for given expression in the specified language
func (l Language) NewEvaluable(expression string) (Evaluable, error) {
	p := newParser(expression)

	eval, err := l.ChainEvaluable(context.Background(), p)
	if err == nil && p.isCamouflaged() && p.lastScan != scanner.EOF {
		err = p.camouflage
	}
	if err != nil {
		pos := p.scanner.Pos()
		return nil, fmt.Errorf("parsing error: %s - %d:%d %s", p.scanner.Position, pos.Line, pos.Column, err)
	}

	return eval, nil
}

// Evaluate given parameter with given expression
func (l Language) Evaluate(expression string, parameter interface{}) (interface{}, error) {
	eval, err := l.NewEvaluable(expression)
	if err != nil {
		return nil, err
	}
	v, err := eval(context.Background(), parameter)
	if err != nil {
		return nil, fmt.Errorf("can not evaluate %s: %v", expression, err)
	}
	return v, nil
}

// Function returns a Language with given function.
// Function has no conversion for input types.
//
// If the function returns an error it must be the last return parameter.
//
// If the function has (without the error) more then one return parameter,
// it returns them as []interface{}.
func Function(name string, function interface{}) Language {
	l := newLanguage()
	l.prefixes[name] = func(c context.Context, p *Parser) (eval Evaluable, err error) {
		args := []Evaluable{}
		scan := p.Scan()
		switch scan {
		case '(':
			args, err = p.parseArguments(c)
			if err != nil {
				return nil, err
			}
		default:
			p.Camouflage("function call", '(')
		}
		return p.callFunc(toFunc(function), args...), nil
	}
	return l
}

// Constant returns a Language with given constant
func Constant(name string, value interface{}) Language {
	l := newLanguage()
	l.prefixes[l.makePrefixKey(name)] = func(c context.Context, p *Parser) (eval Evaluable, err error) {
		return p.Const(value), nil
	}
	return l
}

// PrefixExtension extends a Language
func PrefixExtension(r rune, ext func(context.Context, *Parser) (Evaluable, error)) Language {
	l := newLanguage()
	l.prefixes[r] = ext
	return l
}

// PrefixLanguage switches the parsing language from the current language to the
// given language. When the subordinate language cannot consume any more input
// from the parser, control is returned to the current language and the callback
// is invoked to decide how to continue.
func PrefixLanguage(r rune, sub Language, ext func(context.Context, *Parser, Evaluable) (Evaluable, error)) Language {
	l := newLanguage()
	l.prefixes[r] = func(c context.Context, p *Parser) (Evaluable, error) {
		eval, err := sub.ChainEvaluable(c, p)
		if err != nil {
			return nil, err
		}

		return ext(c, p, eval)
	}
	return l
}

// Init is a language that does no parsing, but invokes the given function when
// parsing starts. It is incumbent upon the function to call ParseExpression to
// continue parsing.
//
// This function can be used to customize the parser settings, such as
// whitespace or ident behavior.
func Init(ext func(context.Context, *Parser) (Evaluable, error)) Language {
	l := newLanguage()
	l.init = ext
	return l
}

// Alternate is a language that runs the given function if no other prefix
// matches.
func Alternate(ext func(context.Context, *Parser) (Evaluable, error)) Language {
	l := newLanguage()
	l.alternate = ext
	return l
}

// Late allows late-binding of a language so that languages can be
// self-referencing. The return value of the passed function will be bound to
// the language.
func Late(fn func(late Language) Language) (binding Language) {
	late := Init(func(ctx context.Context, p *Parser) (Evaluable, error) {
		return binding.ChainEvaluable(ctx, p)
	})
	binding = fn(late)
	return
}

// PrefixMetaPrefix chooses a Prefix to be executed
func PrefixMetaPrefix(r rune, ext func(context.Context, *Parser) (call string, alternative func() (Evaluable, error), err error)) Language {
	l := newLanguage()
	l.prefixes[r] = func(c context.Context, p *Parser) (Evaluable, error) {
		call, alternative, err := ext(c, p)
		if err != nil {
			return nil, err
		}
		if prefix, ok := p.currentLanguage().prefixes[l.makePrefixKey(call)]; ok {
			return prefix(c, p)
		}
		return alternative()
	}
	return l
}

//PrefixOperator returns a Language with given prefix
func PrefixOperator(name string, e Evaluable) Language {
	l := newLanguage()
	l.prefixes[l.makePrefixKey(name)] = func(c context.Context, p *Parser) (Evaluable, error) {
		eval, err := p.ParseNextExpression(c)
		if err != nil {
			return nil, err
		}
		prefix := func(c context.Context, v interface{}) (interface{}, error) {
			a, err := eval(c, v)
			if err != nil {
				return nil, err
			}
			return e(c, a)
		}
		if eval.IsConst() {
			v, err := prefix(context.Background(), nil)
			if err != nil {
				return nil, err
			}
			prefix = p.Const(v)
		}
		return prefix, nil
	}
	return l
}

// PostfixOperator extends a Language.
func PostfixOperator(name string, ext func(context.Context, *Parser, Evaluable) (Evaluable, error)) Language {
	l := newLanguage()
	l.operators[l.makeInfixKey(name)] = postfix{
		f: func(c context.Context, p *Parser, eval Evaluable, pre operatorPrecedence) (Evaluable, error) {
			return ext(c, p, eval)
		},
	}
	return l
}

// InfixOperator for two arbitrary values.
func InfixOperator(name string, f func(a, b interface{}) (interface{}, error)) Language {
	return newLanguageOperator(name, &infix{arbitrary: f})
}

// InfixShortCircuit operator is called after the left operand is evaluated.
func InfixShortCircuit(name string, f func(a interface{}) (interface{}, bool)) Language {
	return newLanguageOperator(name, &infix{shortCircuit: f})
}

// InfixTextOperator for two text values.
func InfixTextOperator(name string, f func(a, b string) (interface{}, error)) Language {
	return newLanguageOperator(name, &infix{text: f})
}

// InfixNumberOperator for two number values.
func InfixNumberOperator(name string, f func(a, b float64) (interface{}, error)) Language {
	return newLanguageOperator(name, &infix{number: f})
}

// InfixBoolOperator for two bool values.
func InfixBoolOperator(name string, f func(a, b bool) (interface{}, error)) Language {
	return newLanguageOperator(name, &infix{boolean: f})
}

// Precedence of operator. The Operator with higher operatorPrecedence is evaluated first.
func Precedence(name string, operatorPrecendence uint8) Language {
	return newLanguageOperator(name, operatorPrecedence(operatorPrecendence))
}

// InfixEvalOperator operates on the raw operands.
// Therefore it cannot be combined with operators for other operand types.
func InfixEvalOperator(name string, f func(a, b Evaluable) (Evaluable, error)) Language {
	return newLanguageOperator(name, directInfix{infixBuilder: f})
}

func newLanguageOperator(name string, op operator) Language {
	op.initiate(name)
	l := newLanguage()
	l.operators[l.makeInfixKey(name)] = op
	return l
}

func (l *Language) makePrefixKey(key string) interface{} {
	runes := []rune(key)
	if len(runes) == 1 && !unicode.IsLetter(runes[0]) {
		return runes[0]
	}
	return key
}

func (l *Language) makeInfixKey(key string) string {
	runes := []rune(key)
	for _, r := range runes {
		l.operatorSymbols[r] = struct{}{}
	}
	return key
}

// VariableSelector returns a Language which uses given variable selector.
// It must be combined with a Language that uses the vatiable selector. E.g. gval.Base().
func VariableSelector(selector func(path Evaluables) Evaluable) Language {
	l := newLanguage()
	l.selector = selector
	return l
}

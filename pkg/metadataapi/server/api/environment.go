package api

import (
	"net/http"

	"github.com/puppetlabs/horsehead/v2/encoding/transfer"
	utilapi "github.com/puppetlabs/horsehead/v2/httputil/api"
	"github.com/puppetlabs/relay-core/pkg/expr/evaluate"
	"github.com/puppetlabs/relay-core/pkg/manager/resolve"
	"github.com/puppetlabs/relay-core/pkg/metadataapi/errors"
	"github.com/puppetlabs/relay-core/pkg/metadataapi/server/middleware"
)

func (s *Server) GetEnvironment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	managers := middleware.Managers(r)

	environment, err := managers.Environment().Get(ctx)
	if err != nil {
		utilapi.WriteError(ctx, w, ModelReadError(err))
		return
	}

	complete := true
	evs := make(map[string]interface{})
	for name, value := range environment.Value {
		ev := evaluate.NewEvaluator(
			evaluate.WithParameterTypeResolver(resolve.NewParameterTypeResolver(managers.Parameters())),
			evaluate.WithOutputTypeResolver(resolve.NewOutputTypeResolver(managers.StepOutputs())),
			evaluate.WithSecretTypeResolver(resolve.NewSecretTypeResolver(managers.Secrets())),
		).ScopeTo(value)

		rv, rerr := ev.EvaluateAll(ctx)
		if rerr != nil {
			utilapi.WriteError(ctx, w, errors.NewExpressionEvaluationError(rerr.Error()))
			return
		}

		if !rv.Complete() {
			complete = false
		}

		evs[name] = rv.Value
	}

	env := &evaluate.JSONResultEnvelope{
		Value:    transfer.JSONInterface{Data: evs},
		Complete: complete,
	}

	utilapi.WriteObjectOK(ctx, w, env)
}

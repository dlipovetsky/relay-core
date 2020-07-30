// Code generated by vfsgen; DO NOT EDIT.

package asset

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// assets statically implements the virtual filesystem provided to vfsgen.
var assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Time{},
		},
		"/schemas": &vfsgen۰DirInfo{
			name:    "schemas",
			modTime: time.Time{},
		},
		"/schemas/v1": &vfsgen۰DirInfo{
			name:    "v1",
			modTime: time.Time{},
		},
		"/schemas/v1/Workflow.json": &vfsgen۰CompressedFileInfo{
			name:             "Workflow.json",
			modTime:          time.Time{},
			uncompressedSize: 6930,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\x4d\x6f\xdb\x38\x13\xbe\xfb\x57\x0c\xd4\x02\xef\xa5\x79\xbd\x3d\x2d\x90\x5b\xda\x6e\x81\x02\x5d\xb4\x68\xda\xed\xa1\xe8\x81\x96\x46\x16\x6b\x8a\x64\x49\xca\x89\x11\xf8\xbf\x2f\x48\x4a\xb2\x28\x51\x8a\x25\x27\xd8\x1c\x92\x48\xe2\xcc\x3c\x7c\xe6\x93\x7c\x58\x01\x24\x2f\x75\x5a\x60\x49\x92\x6b\x48\x0a\x63\xe4\xf5\x7a\xfd\x4b\x0b\x7e\xe5\xdf\xfe\x5f\xa8\xed\x3a\x53\x24\x37\x57\x7f\xfc\xb9\xf6\xef\x5e\x24\xaf\xac\x9c\xa1\x86\xa1\x95\xfa\x2e\xd4\x2e\x67\xe2\xce\xbf\xce\x50\xa7\x8a\x4a\x43\x05\xb7\x1f\x6f\xe0\xae\xfe\x0c\x19\xe6\x94\x53\xf7\xc1\x2b\x38\x48\x27\x2f\x36\xbf\x30\x35\xfe\x9d\x54\x42\xa2\x32\x14\x75\x72\x0d\x16\x1e\x40\x42\x24\xfd\x07\x95\xf6\x0a\xfd\xbb\x8e\xb4\x36\x8a\xf2\xad\x93\x76\xef\x7b\xf6\xbf\x16\x78\x42\xe0\xf1\xc3\xbe\xd6\xd6\xca\x20\xaf\xca\xe4\x1a\x7e\xd4\xcf\x00\xc9\xfe\x75\x52\x3f\xfc\x74\x7f\x8f\x7e\x6d\xb2\xa3\x3c\x7b\x22\x14\x4e\x74\x02\x42\x4b\x6b\x14\x88\xae\xca\x92\xa8\xc3\x02\x2c\x37\x50\xd0\x6d\x71\xc5\x70\x8f\x0c\x64\xa1\x88\x46\xf0\x4b\x36\x94\x6f\xe1\xae\x20\x06\x4c\x17\x6f\x26\x50\x27\x81\xf1\x50\xe3\x7c\x00\xa5\x50\xd6\xa6\x21\x94\x61\x06\x9d\xcf\x20\xf2\xc0\x76\x68\xb6\x10\x25\x4a\xb2\xc5\x25\x36\x39\x08\xf7\x40\x18\x7c\xfb\xf2\x11\x8c\x00\x02\x77\xb8\xd1\xd4\x20\xe4\x42\x85\x56\x5b\x2d\xb9\x50\x25\x31\x56\x41\xa5\x68\x08\x46\x8b\x4a\xa5\x4f\x04\xa5\x6b\xfc\x7f\x1a\x48\x65\x0a\xa1\xa8\x21\x86\xee\x11\xbc\x21\x48\x45\x86\xc0\x44\x4a\x4c\x10\xba\x13\x08\x0d\xd9\xea\x18\x3e\xa2\x14\x39\x9c\x05\x8f\x51\x6d\xac\x4b\x72\x85\x78\x65\x2d\x01\x23\x1b\x64\xda\x62\x2e\x90\x49\x90\x28\x24\x43\xc8\x29\xcf\x46\x18\xa4\x06\xcb\x2e\x8a\x21\x4f\xf5\x87\x63\x80\xdd\x6e\x94\x45\xd1\x77\xaa\x45\x0c\xfe\x17\xac\x34\xd9\x58\x4c\x2d\xe6\x8c\x18\x02\x5a\x62\x4a\x73\x9a\x7a\xba\xa9\x1e\x09\x31\x49\x14\x29\xd1\xa0\x3a\xcb\x36\xc9\x32\xea\xa9\xfa\x3c\xac\x5b\x6e\xc5\x4b\x85\xb9\x95\x7c\xb1\x3e\x55\x3f\xbd\xfe\xdc\x58\x89\x6f\x5e\x1b\x94\x0b\x3c\xf7\xb1\xf6\xd6\xa9\xd0\x38\x3d\xed\xf2\x92\xf2\x0f\xb5\x37\x5e\x4f\xf9\x27\x8e\xf9\xd6\xa0\x8c\xc3\x35\x8a\x6e\xb7\x23\x8c\xcd\x44\xdc\xaa\x9a\x8f\xef\xab\x17\xed\x41\x5c\xd5\x30\x93\xce\xd2\x53\x67\xb9\x75\x95\xf8\x3d\x45\x16\x2d\xea\x8f\x84\xda\x4d\x5b\xb9\x84\x72\x85\xeb\x20\x31\x03\xca\x7d\x4a\x06\x59\x1a\x69\x6b\x5d\x4b\xa7\x37\xe3\x65\x24\x8e\x60\x4f\x58\x85\xbe\x95\xb4\xeb\x8e\x0d\x05\xad\x79\x85\xbf\x2b\xaa\xd0\x6e\xf2\x87\xd7\x1f\x36\x93\xbf\xee\xa5\x42\xdd\xef\xb0\xc3\xc2\x80\xed\x3a\x40\x6b\x99\x18\xcc\x60\x73\x70\xa9\xbf\x21\xe9\x0e\x79\x16\xa6\xd3\x29\xd0\x97\xd0\xdb\x86\x45\x9b\x95\xfd\x19\x62\x92\xdc\x0c\x73\x52\x31\xd3\xe7\xb7\x67\xe6\x9d\x5f\xd5\xb1\xe1\x38\xed\xd0\xf9\x6a\x35\x22\xbb\xd4\x6d\x9f\x9a\xfa\xda\xdd\xd7\x69\xc5\xd0\x91\x01\xa7\x2e\x11\x2f\xa2\xd3\xd6\x85\x59\x4c\x72\x52\x5e\x12\xa6\xdf\x38\xfd\x5d\x61\xcf\xbe\xd3\x39\x42\xb2\x44\x9e\xe9\x4f\x03\x8a\x7b\x6a\x6f\xfd\x36\xec\x62\xe4\xa9\x05\x1d\x80\x10\x1c\x3f\xe5\xc1\x48\x65\x7f\x1e\x82\xa7\xd1\x6e\x34\x40\x35\x29\x1b\x56\xb9\xf6\xf3\xb0\x76\x9d\x65\xf5\xe4\xf8\xe1\xd3\xcf\x33\xb3\xbc\xe7\xbb\x66\x84\x6c\xd7\x0e\xc9\x79\x18\x2b\xac\x6f\x05\x37\x84\x72\x54\x2e\xf0\xba\x9c\x8c\x8a\xdc\x48\xa9\xc4\x9e\xb0\x5a\x22\x3a\xc0\xb6\x6a\xff\xa6\xf7\x34\xa8\x3b\x63\x51\x48\xcb\x70\xf6\x8b\x11\x39\x15\x86\xef\x44\xba\x43\x05\x4e\x8d\x9b\xf9\x5c\x1c\xe2\x3d\xa6\x55\x2f\xef\x3a\xc1\x98\x8a\xb2\x24\xc1\xcc\x3f\xd7\xec\x5b\xaf\xc1\x0e\x1e\x54\xeb\xb1\xca\x42\xd4\xb6\x1f\x2a\x13\xe1\x35\x66\x83\xa8\x6d\x55\x22\x37\xbd\x54\x88\x07\xe2\x44\x10\x1e\xa3\x18\x29\x97\x95\x79\x4f\xd9\x25\x4e\xb0\x27\x21\x85\xcc\x8f\xb6\x92\x98\xc2\x8d\xe2\x1c\x72\xca\xd0\xfe\x5b\xe9\xd3\x3c\xee\xec\x81\x17\x8f\xb3\xe6\x56\x2c\xa7\xed\x43\xc7\x80\x35\xee\x63\x01\x9f\x8e\xbc\x33\x73\xd4\x47\xf6\x74\x9a\xf4\xcb\xfe\x9c\x91\x22\x15\x5c\xbb\x33\x42\xda\x68\x9b\x9a\x16\x08\x63\xb3\x6b\x83\x4f\xe2\x91\x54\x0f\xca\xc1\xa5\x5b\x20\xb5\xb2\x4b\xe6\x9d\x66\x56\xbc\xa8\x8b\xd6\xb3\xea\x7f\xd9\x48\x1b\x08\xe3\xbd\x74\x70\x4a\x85\x47\x27\xe8\x5b\x2f\x12\x55\xb7\xa1\x3c\xb3\x20\xe7\xe8\x7b\x53\xcb\x44\x15\xde\x15\x38\x68\xf3\x71\x6d\x9d\x19\x75\x59\x0f\x1c\xb2\x12\x0d\xd6\x90\x85\x27\x89\x90\xfa\x04\x1f\x0b\x94\x19\x7d\xd8\x1e\x56\xb2\x8a\x61\x08\xf0\xac\x7e\xfc\xb9\xd2\xc5\x02\xb1\xef\xb8\x29\x84\xd8\xf5\x25\xa3\xb4\xc5\xd1\x2d\xa1\x4f\xd7\x9a\xce\xa1\x6f\x51\x05\x69\x0c\x8c\x64\x4c\xf3\xf5\xb2\x1e\x97\xdb\x48\x44\x9e\x1e\x5c\xdf\xe7\x7b\xb1\x43\x7f\xef\xd0\xec\xaa\x3e\x4b\xb9\x13\xa3\xeb\x76\xa9\x12\x1c\xf4\x81\x1b\x72\x3f\x3f\xc2\x83\xcb\x44\x18\xee\x31\x74\xd6\x30\x1e\x96\x38\x4a\x56\xba\x78\x36\x27\x59\xe5\xe3\x0e\x72\xd7\xd5\x51\xf7\xf4\x50\x47\x91\x77\x6e\xb8\x9a\xbb\x58\x01\x68\xcf\xed\x29\x82\x75\x42\xb5\x29\xa9\xb1\x67\x5b\xdc\x23\x37\xee\x02\x29\xd4\xf8\xe8\xc5\x0f\x4c\x5c\xa4\x74\x2e\x1d\x2e\x9a\x1a\xba\xa7\xfe\xd0\xbf\xd1\xc4\x5d\x54\xca\xbc\xa2\x67\xf3\x72\xad\xff\xb9\xa7\x91\xb9\x83\xc1\x9b\x41\x9b\x3b\x9f\xb2\x86\xaa\xba\x55\xce\xe2\x6a\x87\x87\x4b\x6e\x83\x3a\x71\xbd\xc3\x43\x8b\xc0\x4e\xd5\x6e\x70\x60\x07\xa0\x19\x72\x43\xf3\x83\x1d\xb9\x5d\x70\xc7\x73\x2c\x7a\x0d\x3a\xc5\x43\x0c\x4f\x7b\xf9\xd3\x20\xd1\xcd\x80\x6f\xbb\x7e\x43\x94\xbf\x41\x6a\xe1\x3c\x59\x9a\xc5\xe6\x85\x68\x96\xad\x9a\xdf\xc7\xd5\x71\xf5\x6f\x00\x00\x00\xff\xff\x98\x35\x1b\x0d\x12\x1b\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/schemas"].(os.FileInfo),
	}
	fs["/schemas"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/schemas/v1"].(os.FileInfo),
	}
	fs["/schemas/v1"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/schemas/v1/Workflow.json"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}

//js:package fs

package fs

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/gosolid/solid/pkg/runtime/buffer"
	"github.com/grexie/isolates"
)

var _ File = &httpfile{}
var _ filev8 = &httpfile{}
var _ FS = &httpfs{}
var _ fsv8 = &httpfs{}

type httpfile struct {
	*filebase
	path string
}

type httpfs struct {
	*fs
	url url.URL
}

//js:constructor HttpFileSystem
//js:export HttpFileSystem
func NewHttpFS(url url.URL) *httpfs {
	return &httpfs{NewFS(), url}
}

//js:method toString
func (fs *httpfs) String() string {
	return fmt.Sprintf("[HttpFS %s]", fs.url)
}

func (fs *httpfs) joinPath(p string) (string, error) {
	u := fs.url
	u.Path = path.Join(u.Path, p)
	if !strings.HasPrefix(u.String(), fs.url.String()) {
		return "", fmt.Errorf("file not found")
	}
	return u.String(), nil
}

//js:method readdirSync
//js:callback-method readdir
func (fs *httpfs) ReadDir(ctx context.Context, path string) ([]string, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && err != os.ErrNotExist {
		return nil, err
	} else if err != os.ErrNotExist {
		return mount.fs.ReadDir(ctx, mount.path)
	}

	return nil, fmt.Errorf("not supported")
}

//js:method readFileSync
//js:callback-method readFile
//js:return buffer.Buffer
func (fs *httpfs) ReadFile(ctx context.Context, path string, options ...any) (any, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && err != os.ErrNotExist {
		return nil, err
	} else if err != os.ErrNotExist {
		return mount.fs.ReadFile(ctx, mount.path, options...)
	}

	if encoding, err := parseReadFileOptions(ctx, options); err != nil {
		return nil, err
	} else if file, err := fs.Open(ctx, path); err != nil {
		return nil, err
	} else if bytes, err := file.ReadAll(ctx); err != nil {
		return nil, err
	} else {
		if encoding != nil {
			if b, err := isolates.For(ctx).New(buffer.NewBuffer, len(bytes)); err != nil {
				return nil, err
			} else if err := b.(*buffer.Buffer).Buffer().SetBytes(ctx, bytes); err != nil {
				return nil, err
			} else if s, err := b.(*buffer.Buffer).ToString(ctx, encoding); err != nil {
				return nil, err
			} else {
				return s, nil
			}
		}
		return bytes, nil

	}
}

//js:method createReadStream
func (fs *httpfs) ReadStream(ctx context.Context, path string) (ReadStream, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && err != os.ErrNotExist {
		return nil, err
	} else if err != os.ErrNotExist {
		return mount.fs.ReadStream(ctx, mount.path)
	}

	if file, err := fs.Open(ctx, path); err != nil {
		return nil, err
	} else if stream, err := file.ReadStream(ctx); err != nil {
		return nil, err
	} else {
		return stream, nil
	}
}

//js:method openSync
//js:callback-method open
func (fs *httpfs) Open(ctx context.Context, path string) (FileDescriptor, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && !IsNotExists(err) {
		return 0, err
	} else if !IsNotExists(err) {
		return mount.fs.Open(ctx, mount.path)
	}

	if u, err := fs.joinPath(path); err != nil {
		return 0, err
	} else if req, err := http.NewRequest(http.MethodHead, u, nil); err != nil {
		return 0, err
	} else if res, err := http.DefaultClient.Do(req); err != nil {
		return 0, err
	} else if res.StatusCode != 200 {
		return 0, os.ErrNotExist
	} else {
		return NewFileDescriptor(ctx, &httpfile{&filebase{}, u})
	}
}

//js:method closeSync
//js:callback-method close
func (fs *httpfs) Close(ctx context.Context, fd FileDescriptor) error {
	return fd.Close(ctx)
}

//js:method file
func (fs *httpfs) File(ctx context.Context, fd FileDescriptor) (File, error) {
	return fd.File(ctx)
}

//js:method existsSync
//js:callback-method exists
func (fs *httpfs) Exists(ctx context.Context, p string) (bool, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, p); err != nil && !IsNotExists(err) {
		return false, err
	} else if !IsNotExists(err) {
		return mount.fs.Exists(ctx, mount.path)
	}

	if _, err := fs.Stat(ctx, p); err != nil && !IsNotExists(err) {
		return false, err
	} else {
		return err == nil || !IsNotExists(err), nil
	}
}

//js:method statSync
//js:callback-method stat
func (fs *httpfs) Stat(ctx context.Context, path string) (Stats, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && err != os.ErrNotExist {
		return nil, err
	} else if err != os.ErrNotExist {
		return mount.fs.Stat(ctx, mount.path)
	}

	if u, err := fs.joinPath(path); err != nil {
		return nil, err
	} else if req, err := http.NewRequest(http.MethodHead, u, nil); err != nil {
		return nil, err
	} else if res, err := http.DefaultClient.Do(req); err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, os.ErrNotExist
	} else {
		return &stats{
			isFile:      true,
			isDirectory: false,
			contentType: res.Header.Get("Content-Type"),
			modified:    0,
			size:        0,
		}, nil
	}
}

//js:method lstatSync
//js:callback-method lstat
func (fs *httpfs) LStat(ctx context.Context, p string) (Stats, error) {
	return fs.Stat(ctx, p)
}

//js:callback-method realpath
//js:method realpathSync
func (f *httpfs) RealPath(ctx context.Context, p string) (string, error) {
	if mount, mountPoint, err := f.mounts.Resolve(ctx, p); err != nil && !os.IsNotExist(err) {
		return "", err
	} else if !os.IsNotExist(err) {
		if realpath, err := mount.fs.RealPath(ctx, mount.path); err != nil {
			return "", err
		} else {
			return path.Join(mountPoint, realpath), nil
		}
	}

	return path.Clean(p), nil
}

//js:method closeSync
//js:callback-method close
func (f *httpfile) Close(ctx context.Context) error {
	return nil
}

func (f *httpfile) ReadStream(ctx context.Context) (ReadStream, error) {
	if req, err := http.NewRequest(http.MethodGet, f.path, nil); err != nil {
		return nil, err
	} else if res, err := http.DefaultClient.Do(req); err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, os.ErrNotExist
	} else if stream, err := isolates.For(ctx).New(NewReadStream, res.Body); err != nil {
		return nil, err
	} else {
		return stream.(ReadStream), nil
	}
}

//js:method readAllSync
//js:callback-method readAll
//js:return buffer.Buffer
func (f *httpfile) ReadAll(ctx context.Context) ([]byte, error) {
	if reader, err := f.ReadStream(ctx); err != nil {
		return nil, err
	} else {
		return io.ReadAll(reader)
	}
}

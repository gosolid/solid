//js:package fs

package fs

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"fmt"
	"io"
	"mime"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gosolid/solid/pkg/runtime/buffer"
	isolates "github.com/grexie/isolates"
)

var _ File = &localfile{}
var _ filev8 = &localfile{}
var _ FS = &localfs{}
var _ fsv8 = &localfs{}

type localfile struct {
	*filebase
	fs   *localfs
	file *os.File
}

type localfs struct {
	*fs
	path string
}

//js:constructor LocalFileSystem
//js:export LocalFileSystem
func NewLocalFS(path string) *localfs {
	return &localfs{NewFS(), path}
}

//js:method toString
func (fs *localfs) String() string {
	return fmt.Sprintf("[fs Local %s]", fs.path)
}

func (fs *localfs) joinPath(p string) (string, error) {
	p = path.Join(fs.path, p)
	if !strings.HasPrefix(p, fs.path) {
		return "", fmt.Errorf("file not found")
	}
	return p, nil
}

//js:method
func (fs *localfs) Mount(mountPoint string, mount FS, path string) error {
	fs.mounts[mountPoint] = &locator{mount, path}
	return nil
}

//js:method readdirSync
//js:callback-method readdir
func (fs *localfs) ReadDir(ctx context.Context, path string) ([]string, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && !IsNotExists(err) {
		return nil, err
	} else if !IsNotExists(err) {
		return mount.fs.ReadDir(ctx, mount.path)
	}

	if p, err := fs.joinPath(path); err != nil {
		return nil, err
	} else if entries, err := os.ReadDir(p); err != nil {
		if os.IsNotExist(err) {
			return nil, NewFSError(ctx, ENOENT, err.Error())
		} else {
			return nil, err
		}
	} else {
		files := []string{}
		for _, entry := range entries {
			files = append(files, entry.Name())
		}
		return files, nil
	}
}

//js:method readFileSync
//js:callback-method readFile
//js:return buffer.Buffer
func (fs *localfs) ReadFile(ctx context.Context, path string, options ...any) (any, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && !IsNotExists(err) {
		return nil, err
	} else if !IsNotExists(err) {
		return mount.fs.ReadFile(ctx, mount.path, options...)
	}

	if encoding, err := parseReadFileOptions(ctx, options); err != nil {
		return nil, err
	} else if p, err := fs.joinPath(path); err != nil {
		return nil, err
	} else if bytes, err := os.ReadFile(p); err != nil {
		if os.IsNotExist(err) {
			return nil, NewFSError(ctx, ENOENT, err.Error())
		} else {
			return nil, err
		}
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
func (fs *localfs) ReadStream(ctx context.Context, path string) (ReadStream, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && !IsNotExists(err) {
		return nil, err
	} else if !IsNotExists(err) {
		return mount.fs.ReadStream(ctx, mount.path)
	}

	if f, err := fs.Open(ctx, path); err != nil {
		return nil, err
	} else {
		return f.ReadStream(ctx)
	}
}

//js:method openSync
//js:callback-method open
func (fs *localfs) Open(ctx context.Context, path string) (FileDescriptor, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && !IsNotExists(err) {
		return 0, err
	} else if !IsNotExists(err) {
		return mount.fs.Open(ctx, mount.path)
	}

	if p, err := fs.joinPath(path); err != nil {
		return 0, err
	} else if file, err := os.OpenFile(p, os.O_RDONLY, 0); err != nil {
		if os.IsNotExist(err) {
			return 0, NewFSError(ctx, ENOENT, err.Error())
		} else {
			return 0, err
		}
	} else {
		return NewFileDescriptor(ctx, &localfile{&filebase{}, fs, file})
	}
}

//js:method existsSync
//js:callback-method exists
func (fs *localfs) Exists(ctx context.Context, p string) (bool, error) {
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
func (fs *localfs) Stat(ctx context.Context, p string) (Stats, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, p); err != nil && !IsNotExists(err) {
		return nil, err
	} else if !IsNotExists(err) {
		return mount.fs.Stat(ctx, mount.path)
	}

	if p, err := fs.joinPath(p); err != nil {
		return nil, err
	} else if s, err := os.Stat(p); err != nil {
		if os.IsNotExist(err) {
			return nil, NewFSError(ctx, ENOENT, err.Error())
		} else {
			return nil, err
		}
	} else {
		return &stats{
			isFile:      !s.IsDir(),
			isDirectory: s.IsDir(),
			contentType: mime.TypeByExtension(path.Ext(p)),
			modified:    s.ModTime().UnixMilli(),
			size:        s.Size(),
		}, nil
	}
}

//js:method lstatSync
//js:callback-method lstat
func (fs *localfs) LStat(ctx context.Context, p string) (Stats, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, p); err != nil && !IsNotExists(err) {
		return nil, err
	} else if !IsNotExists(err) {
		return mount.fs.LStat(ctx, mount.path)
	}

	if p, err := fs.joinPath(p); err != nil {
		return nil, err
	} else if s, err := os.Lstat(p); err != nil {
		if os.IsNotExist(err) {
			return nil, NewFSError(ctx, ENOENT, err.Error())
		} else {
			return nil, err
		}
	} else {
		return &stats{
			isFile:      !s.IsDir(),
			isDirectory: s.IsDir(),
			contentType: mime.TypeByExtension(path.Ext(p)),
			modified:    s.ModTime().UnixMilli(),
			size:        s.Size(),
		}, nil
	}
}

//js:callback-method realpath
//js:method realpathSync
func (f *localfs) RealPath(ctx context.Context, p string) (string, error) {
	if mount, _, err := f.mounts.Resolve(ctx, p); err != nil && !IsNotExists(err) {
		return "", err
	} else if !IsNotExists(err) {
		return mount.fs.RealPath(ctx, mount.path)
	}

	if p2, err := f.joinPath(p); err != nil {
		return "", err
	} else if realpath, err := filepath.EvalSymlinks(p2); err != nil {
		return "", err
	} else if rel, err := filepath.Rel(f.path, realpath); err != nil {
		return "", err
	} else {
		if strings.HasPrefix(rel, "../") {
			return path.Clean(p), nil
		} else {
			if rel, err := filepath.Rel(f.path, path.Clean(realpath)); err != nil {
				return "", err
			} else {
				return "/" + rel, nil
			}
		}
	}
}

//js:method closeSync
//js:callback-method
func (f *localfile) Close(ctx context.Context) error {
	return f.file.Close()
}

func (f *localfile) ReadStream(ctx context.Context) (ReadStream, error) {
	if stream, err := isolates.For(ctx).New(NewReadStream, f.file); err != nil {
		return nil, err
	} else {
		return stream.(ReadStream), nil
	}
}

//js:method readAllSync
//js:callback-method readAll
//js:return buffer.Buffer
func (f *localfile) ReadAll(ctx context.Context) ([]byte, error) {
	if reader, err := f.ReadStream(ctx); err != nil {
		return nil, err
	} else {
		return io.ReadAll(reader)
	}
}

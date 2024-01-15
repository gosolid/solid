//js:package fs

package fs

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"embed"
	"fmt"
	"io"
	gofs "io/fs"
	"mime"
	"path"

	"github.com/gosolid/solid/pkg/runtime/buffer"
	"github.com/grexie/isolates"
)

var _ buffer.Buffer

var _ File = &embedfile{}
var _ filev8 = &embedfile{}
var _ FS = &embedfs{}
var _ fsv8 = &embedfs{}

type embedfile struct {
	*filebase
	fs   *embedfs
	file gofs.File
}

type embedfs struct {
	*fs
	embed *embed.FS
}

//js:constructor EmbeddedFileSystem
//js:export EmbeddedFileSystem
func NewEmbedFS(embed *embed.FS) *embedfs {
	return &embedfs{NewFS(), embed}
}

//js:method toString
func (fs *embedfs) String() string {
	return fmt.Sprintf("[fs Embedded]")
}

//js:method createReadStream
func (fs *embedfs) ReadStream(ctx context.Context, path string) (ReadStream, error) {
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

//js:method readdirSync
//js:callback-method readdir
func (fs *embedfs) ReadDir(ctx context.Context, path string) ([]string, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && !IsNotExists(err) {
		return nil, err
	} else if !IsNotExists(err) {
		return mount.fs.ReadDir(ctx, mount.path)
	}

	if entries, err := fs.embed.ReadDir(path[1:]); err != nil {
		return nil, err
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
func (fs *embedfs) ReadFile(ctx context.Context, path string) ([]byte, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && !IsNotExists(err) {
		return nil, err
	} else if !IsNotExists(err) {
		return mount.fs.ReadFile(ctx, mount.path)
	}

	if bytes, err := fs.embed.ReadFile(path[1:]); err != nil {
		return nil, err
	} else {
		return bytes, nil
	}
}

//js:method openSync
//js:callback-method open
func (fs *embedfs) Open(ctx context.Context, path string) (FileDescriptor, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && !IsNotExists(err) {
		return 0, err
	} else if !IsNotExists(err) {
		return mount.fs.Open(ctx, mount.path)
	}

	if file, err := fs.embed.Open(path[1:]); err != nil {
		return 0, err
	} else {

		return NewFileDescriptor(ctx, &embedfile{&filebase{}, fs, file})
	}
}

//js:method closeSync
//js:callback-method close
func (fs *embedfs) Close(ctx context.Context, fd FileDescriptor) error {
	return fd.Close(ctx)
}

//js:method file
func (fs *embedfs) File(ctx context.Context, fd FileDescriptor) (File, error) {
	return fd.File(ctx)
}

//js:method existsSync
//js:callback-method exists
func (fs *embedfs) Exists(ctx context.Context, p string) (bool, error) {
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
func (fs *embedfs) Stat(ctx context.Context, p string) (Stats, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, p); err != nil && !IsNotExists(err) {
		return nil, err
	} else if !IsNotExists(err) {
		return mount.fs.Stat(ctx, mount.path)
	}

	if _, err := fs.embed.ReadDir(p[1:]); err == nil {
		return &stats{
			isFile:      false,
			isDirectory: true,
			contentType: "application/directory",
			modified:    0,
			size:        0,
		}, nil
	} else if file, err := fs.embed.ReadFile(p[1:]); err == nil {
		return &stats{
			isFile:      true,
			isDirectory: false,
			contentType: mime.TypeByExtension(path.Ext(p)),
			modified:    0,
			size:        int64(len(file)),
		}, nil
	} else {
		return nil, fmt.Errorf("ENOENT: file not found")
	}
}

//js:method lstatSync
//js:callback-method lstat
func (fs *embedfs) LStat(ctx context.Context, p string) (Stats, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, p); err != nil && !IsNotExists(err) {
		return nil, err
	} else if !IsNotExists(err) {
		return mount.fs.LStat(ctx, mount.path)
	}

	return fs.Stat(ctx, p)
}

//js:callback-method realpath
//js:method realpathSync
func (f *embedfs) RealPath(ctx context.Context, p string) (string, error) {
	if path, err := f.fs.RealPath(ctx, p); err == nil || !IsNotExists(err) {
		return path, err
	}

	return path.Clean(p), nil
}

//js:callback-method close
//js:method closeSync
func (f *embedfile) Close(ctx context.Context) error {
	return f.file.Close()
}

func (f *embedfile) ReadStream(ctx context.Context) (ReadStream, error) {
	if stream, err := isolates.For(ctx).New(NewReadStream, f.file); err != nil {
		return nil, err
	} else {
		return stream.(ReadStream), nil
	}
}

//js:method readAllSync
//js:callback-method readAll
//js:return buffer.Buffer
func (f *embedfile) ReadAll(ctx context.Context) ([]byte, error) {
	if reader, err := f.ReadStream(ctx); err != nil {
		return nil, err
	} else {
		return io.ReadAll(reader)
	}
}

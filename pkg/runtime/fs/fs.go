//js:package fs

package fs

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"fmt"
	gofs "io/fs"
	"os"

	"github.com/gosolid/solid/pkg/runtime/buffer"
	isolates "github.com/grexie/isolates"
)

var _ buffer.Buffer

var _ FS = &fs{}
var _ fsv8 = &fs{}

type fs struct {
	mounts mounts
}

//js:class File
//js:export File
type filebase struct {
}

//js:constructor FileSystem
//js:export FileSystem
func NewFS() *fs {
	return &fs{mounts{}}
}

func (f *fs) V8Construct(in isolates.FunctionArgs) {
	in.This.RebindAll(in.ExecutionContext)
}

//js:method toString
func (fs *fs) String() string {
	return fmt.Sprintf("[fs Empty]")
}

//js:method
func (fs *fs) Mount(mountPoint string, mount FS, path string) error {
	fs.mounts[mountPoint] = &locator{mount, path}
	return nil
}

//js:method createReadStream
func (fs *fs) ReadStream(ctx context.Context, path string) (ReadStream, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && !IsNotExists(err) {
		return nil, err
	} else if !IsNotExists(err) {
		return mount.fs.ReadStream(ctx, mount.path)
	}

	return nil, NewFSError(ctx, ENOENT, fmt.Sprintf("%s does not exist", path))
}

//js:method readdirSync
//js:callback-method readdir
func (fs *fs) ReadDir(ctx context.Context, path string) ([]string, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && !IsNotExists(err) {
		return nil, err
	} else if !IsNotExists(err) {
		return mount.fs.ReadDir(ctx, mount.path)
	}

	return nil, NewFSError(ctx, ENOENT, fmt.Sprintf("%s does not exist", path))
}

//js:callback-method realpath
//js:method realpathSync
func (f *fs) RealPath(ctx context.Context, p string) (string, error) {
	if mount, _, err := f.mounts.Resolve(ctx, p); err != nil && !IsNotExists(err) {
		return "", err
	} else if !IsNotExists(err) {
		if realpath, err := mount.fs.RealPath(ctx, mount.path); err != nil {
			return "", err
		} else {
			return mount.Reverse(realpath), nil
		}
	}

	return "", NewFSError(ctx, ENOENT, fmt.Sprintf("%s does not exist", p))
}

//js:method readFileSync
//js:callback-method readFile
//js:return buffer.Buffer
func (fs *fs) ReadFile(ctx context.Context, path string) ([]byte, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && !IsNotExists(err) {
		return nil, err
	} else if !IsNotExists(err) {
		return mount.fs.ReadFile(ctx, mount.path)
	}

	return nil, NewFSError(ctx, ENOENT, fmt.Sprintf("%s does not exist", path))
}

//js:method openSync
//js:callback-method open
func (fs *fs) Open(ctx context.Context, path string) (FileDescriptor, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, path); err != nil && !IsNotExists(err) {
		return 0, err
	} else if !IsNotExists(err) {
		return mount.fs.Open(ctx, mount.path)
	}

	return 0, NewFSError(ctx, ENOENT, fmt.Sprintf("%s does not exist", path))
}

//js:method closeSync
//js:callback-method close
func (fs *fs) Close(ctx context.Context, fd FileDescriptor) error {
	return fd.Close(ctx)
}

//js:method file
func (fs *fs) File(ctx context.Context, fd FileDescriptor) (File, error) {
	return fd.File(ctx)
}

//js:method existsSync
//js:callback-method exists
func (fs *fs) Exists(ctx context.Context, p string) (bool, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, p); err != nil && !IsNotExists(err) {
		return false, err
	} else if !IsNotExists(err) {
		return mount.fs.Exists(ctx, mount.path)
	}

	return false, nil
}

//js:method statSync
//js:callback-method stat
func (fs *fs) Stat(ctx context.Context, p string) (Stats, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, p); err != nil && !IsNotExists(err) {
		return nil, err
	} else if !IsNotExists(err) {
		return mount.fs.Stat(ctx, mount.path)
	}

	return nil, NewFSError(ctx, ENOENT, fmt.Sprintf("%s does not exist", p))
}

//js:method lstatSync
//js:callback-method lstat
func (fs *fs) LStat(ctx context.Context, p string) (Stats, error) {
	if mount, _, err := fs.mounts.Resolve(ctx, p); err != nil && !IsNotExists(err) {
		return nil, err
	} else if !IsNotExists(err) {
		return mount.fs.LStat(ctx, mount.path)
	}

	return nil, NewFSError(ctx, ENOENT, fmt.Sprintf("%s does not exist", p))
}

//js:method accessSync
//js:callback-method access
func (fs *fs) Access(ctx context.Context, p string, args ...any) error {
	if mount, _, err := fs.mounts.Resolve(ctx, p); err != nil && !IsNotExists(err) {
		return err
	} else if !IsNotExists(err) {
		return mount.fs.Access(ctx, mount.path)
	}

	var mode = R_OK | W_OK

	if args, err := isolates.For(ctx).Context().CreateAll(ctx, args...); err != nil {
		return err
	} else if groups, err := os.Getgroups(); err != nil {
		return err
	} else {
		stats, _ := fs.Stat(ctx, p)
		if len(args) > 1 {
			if m, err := args[0].Int64(ctx); err != nil {
				return err
			} else {
				mode = int(m)
			}
		}

		perm := gofs.FileMode(stats.Mode()).Perm()
		uid := os.Getuid()
		gid := -1

		for _, group := range groups {
			if group == stats.Gid() {
				gid = group
			}
		}

		if mode&F_OK == F_OK {
			if stats == nil {
				return NewFSError(ctx, ENOACCESS, fmt.Sprintf("%s cannot be accessed by current user", p))
			}
		}

		if mode&R_OK == R_OK {
			if !((uid == 0 || stats.Uid() == uid) && perm&0100 == 0100 || (gid == 0 || stats.Gid() == gid) && perm&0010 == 0010 || perm&0001 == 0001) {
				return NewFSError(ctx, ENOACCESS, fmt.Sprintf("%s cannot be read by current user", p))
			}
		}

		if mode&W_OK == W_OK {
			if !((uid == 0 || stats.Uid() == uid) && perm&0200 == 0200 || (gid == 0 || stats.Gid() == gid) && perm&0020 == 0020 || perm&0002 == 0002) {
				return NewFSError(ctx, ENOACCESS, fmt.Sprintf("%s cannot be written by current user", p))
			}
		}

		if mode&X_OK == X_OK {
			if !((uid == 0 || stats.Uid() == uid) && perm&0400 == 0400 || (gid == 0 || stats.Gid() == gid) && perm&0040 == 0040 || perm&0004 == 0004) {
				return NewFSError(ctx, ENOACCESS, fmt.Sprintf("%s cannot be executed by current user", p))
			}
		}

		return nil
	}
}

//js:callback-method close
//js:method closeSync
func (f *filebase) Close(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}

//js:method readAllSync
//js:callback-method readAll
//js:return buffer.Buffer
func (f *filebase) ReadAll(ctx context.Context) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

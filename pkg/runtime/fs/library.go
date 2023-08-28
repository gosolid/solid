//js:package fs

package fs

import (
	"context"
	gofs "io/fs"
	"time"

	"github.com/grexie/isolates"
)

const (
	F_OK = 1
	R_OK = 2
	W_OK = 4
	X_OK = 8
)

type FileMode gofs.FileMode

//go:generate go run github.com/grexie/isolates/codegen

type FS interface {
	ReadDir(ctx context.Context, path string) ([]string, error)
	ReadFile(ctx context.Context, path string) ([]byte, error)
	ReadStream(ctx context.Context, path string) (ReadStream, error)
	Open(ctx context.Context, path string) (FileDescriptor, error)
	Close(ctx context.Context, fd FileDescriptor) error
	Exists(ctx context.Context, path string) (bool, error)
	Stat(ctx context.Context, path string) (Stats, error)
	LStat(ctx context.Context, path string) (Stats, error)
	Access(ctx context.Context, path string, args ...any) error
	Mount(mountPoint string, mount FS, path string) error
	RealPath(ctx context.Context, path string) (string, error)
	File(ctx context.Context, fd FileDescriptor) (File, error)
}

type fsv8 interface {
	V8FuncMount(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncReaddir(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncReaddirSync(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncReadFile(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncReadFileSync(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncCreateReadStream(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncOpen(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncOpenSync(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncClose(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncCloseSync(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncExists(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncExistsSync(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncStat(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncStatSync(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncLstat(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncLstatSync(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncAccess(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncAccessSync(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncRealpath(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncRealpathSync(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncFile(in isolates.FunctionArgs) (*isolates.Value, error)
}

type Stats interface {
	IsFile() bool
	IsDirectory() bool
	IsSymbolicLink() bool
	AccessedMilli() int64
	Accessed() time.Time
	ModifiedMilli() int64
	Modified() time.Time
	ChangedMilli() int64
	Changed() time.Time
	CreatedMilli() int64
	Created() time.Time
	Size() int64
	ContentType() string
	Uid() int
	Gid() int
	Mode() FileMode
}

type statsv8 interface {
	V8FuncIsFile(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncIsDirectory(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncIsSymbolicLink(in isolates.FunctionArgs) (*isolates.Value, error)
	V8GetAtime(in isolates.GetterArgs) (*isolates.Value, error)
	V8GetAtimeMs(in isolates.GetterArgs) (*isolates.Value, error)
	V8GetMtime(in isolates.GetterArgs) (*isolates.Value, error)
	V8GetMtimeMs(in isolates.GetterArgs) (*isolates.Value, error)
	V8GetCtime(in isolates.GetterArgs) (*isolates.Value, error)
	V8GetCtimeMs(in isolates.GetterArgs) (*isolates.Value, error)
	V8GetBirthtime(in isolates.GetterArgs) (*isolates.Value, error)
	V8GetBirthtimeMs(in isolates.GetterArgs) (*isolates.Value, error)
	V8GetSize(in isolates.GetterArgs) (*isolates.Value, error)
	V8GetUid(in isolates.GetterArgs) (*isolates.Value, error)
	V8GetGid(in isolates.GetterArgs) (*isolates.Value, error)
	V8GetMode(in isolates.GetterArgs) (*isolates.Value, error)
}

type DirEntry interface {
	IsFile() bool
	IsDirectory() bool
	Name() string
}

type direntryv8 interface {
	V8FuncIsFile(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncIsDirectory(in isolates.FunctionArgs) (*isolates.Value, error)
	V8GetName(in isolates.GetterArgs) (*isolates.Value, error)
}

type File interface {
	Close(ctx context.Context) error
	ReadStream(ctx context.Context) (ReadStream, error)
	ReadAll(ctx context.Context) ([]byte, error)
}

type filev8 interface {
	V8FuncClose(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncCloseSync(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncCreateReadStream(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncReadAll(in isolates.FunctionArgs) (*isolates.Value, error)
	V8FuncReadAllSync(in isolates.FunctionArgs) (*isolates.Value, error)
}

//js:method createLibrary
//js:export-instance default
func createLibrary(in isolates.RuntimeFunctionArgs) (*isolates.Value, error) {
	if pkg, err := in.Require.Call(in.ExecutionContext, nil, "task"); err != nil {
		return nil, err
	} else if task, err := pkg.Get(in.ExecutionContext, "task"); err != nil {
		return nil, err
	} else if fs, err := task.Get(in.ExecutionContext, "fs"); err != nil {
		return nil, err
	} else {
		return fs, nil
	}
}

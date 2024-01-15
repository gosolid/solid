//js:package fs

package fs

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"sync"
	"unsafe"

	"github.com/grexie/isolates"
)

const kFileDescriptorTable = "solid.fs.fileDescriptorTable"

var ErrNoFileDescriptorTable = func(ctx context.Context) error {
	return NewFSError(ctx, "ENOCTX", "file descriptor table not found in context")
}
var ErrFileDescriptorNotFound = func(ctx context.Context) error { return NewFSError(ctx, "ENOFD", "file descriptor not found") }
var ErrInvalid = func(ctx context.Context) error { return NewFSError(ctx, "EINVAL", "received invalid argument") }

//ts:export type FileDescriptor = number
type FileDescriptor uintptr

func NewFileDescriptor(ctx context.Context, file File) (FileDescriptor, error) {
	if file == nil {
		return 0, ErrInvalid(ctx)
	}

	fd := FileDescriptor(unsafe.Pointer(&file))

	if table, ok := isolates.For(ctx).Isolate().GetData(kFileDescriptorTable).(*sync.Map); !ok {
		return 0, ErrNoFileDescriptorTable(ctx)
	} else {
		table.Store(fd, file)
		return fd, nil
	}
}

func (f FileDescriptor) File(ctx context.Context) (File, error) {
	if table, ok := isolates.For(ctx).Isolate().GetData(kFileDescriptorTable).(*sync.Map); !ok {
		return nil, ErrNoFileDescriptorTable(ctx)
	} else if file, ok := table.Load(f); !ok {
		return nil, ErrFileDescriptorNotFound(ctx)
	} else {
		return file.(File), nil
	}
}

func (f FileDescriptor) Close(ctx context.Context) error {
	if file, err := f.File(ctx); err != nil {
		return err
	} else {
		return file.Close(ctx)
	}
}

func (f FileDescriptor) ReadStream(ctx context.Context) (ReadStream, error) {
	if file, err := f.File(ctx); err != nil {
		return nil, err
	} else {
		return file.ReadStream(ctx)
	}
}

func (f FileDescriptor) ReadAll(ctx context.Context) ([]byte, error) {
	if file, err := f.File(ctx); err != nil {
		return nil, err
	} else {
		return file.ReadAll(ctx)
	}
}

var _ = isolates.RegisterRuntime("fs", "fs/fd.go", func(in isolates.FunctionArgs) (*isolates.Value, error) {
	var table sync.Map

	in.Context.GetIsolate().SetData(kFileDescriptorTable, &table)

	return nil, nil
})

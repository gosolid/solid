//js:package fs

package fs

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"fmt"
	"os"

	"github.com/grexie/isolates"
)

const (
	ENOENT    = "ENOENT"
	ENOACCESS = "ENOACCESS"
)

func IsNotExists(err error) bool {
	if err, ok := err.(FSError); !ok {
		return os.IsNotExist(err)
	} else {
		return err.Code() == ENOENT
	}
}

type FSError interface {
	isolates.Error
	Code() string
}

type FSErrorBase struct {
	code    string
	message string
}

//js:constructor FSError
//js:export FSError
func NewFSError2(in isolates.FunctionArgs) (*FSErrorBase, error) {
	if code, err := in.Arg(in.ExecutionContext, 0).StringValue(in.ExecutionContext); err != nil {
		return nil, err
	} else if message, err := in.Arg(in.ExecutionContext, 1).StringValue(in.ExecutionContext); err != nil {
		return nil, err
	} else {
		return &FSErrorBase{code, message}, nil
	}
}

func NewFSError(ctx context.Context, code string, message string) *FSErrorBase {
	err, _ := isolates.For(ctx).New(NewFSError2, code, message)
	return err.(*FSErrorBase)
}

func (e *FSErrorBase) Error() string {
	return fmt.Sprintf("%s: %s", e.code, e.message)
}

//js:get
func (e *FSErrorBase) Code() string {
	return e.code
}

//js:get
func (e *FSErrorBase) Message() string {
	return e.message
}

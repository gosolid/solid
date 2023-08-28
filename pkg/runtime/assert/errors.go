//js:package assert

package assert

import (
	"context"
	"fmt"

	"github.com/grexie/isolates"
)

const (
	AssertionErrorName string = "AssertionError"
	AssertionErrorCode        = "ERR_ASSERTION"
)

func IsAssertionError(err error) bool {
	_, ok := err.(AssertionError)
	return ok
}

type AssertionError interface {
	isolates.Error
	Actual() *isolates.Value
	Expected() *isolates.Value
	Operator() *string
	GeneratedMessage() bool
}

type AssertionErrorBase struct {
	context *isolates.Context

	message  *string
	actual   *isolates.Value
	expected *isolates.Value
	operator *string
}

//js:constructor AssertionError
//js:export AssertionError
func newAssertionError(ctx context.Context, options AssertionErrorOptions) (*AssertionErrorBase, error) {
	if actualv, err := isolates.For(ctx).Context().Create(ctx, options.actual); err != nil {
		return nil, err
	} else if expectedv, err := isolates.For(ctx).Context().Create(ctx, options.expected); err != nil {
		return nil, err
	} else {
		return &AssertionErrorBase{
			context:  isolates.For(ctx).Context(),
			message:  options.message,
			operator: options.operator,
			actual:   actualv,
			expected: expectedv,
		}, nil
	}
}

type AssertionErrorOptions struct {
	message  *string
	actual   any
	expected any
	operator *string
}

func (e *AssertionErrorBase) Error() string {
	return e.String()
}

//js:method toString
func (e *AssertionErrorBase) String() string {
	diff := ""
	// ctx := e.context.GetIsolate().GetExecutionContext()

	return fmt.Sprintf("%s [%s]: %s%s", AssertionErrorName, AssertionErrorCode, e.Message(), diff)
}

//js:get
func (e *AssertionErrorBase) Message() string {
	msg := ""

	if e.message == nil {
		msg = "generated message"
	} else {
		msg = *e.message
	}

	return msg
}

//js:get
func (e *AssertionErrorBase) Name() string {
	return AssertionErrorName
}

//js:get
func (e *AssertionErrorBase) Actual() any {
	return e.actual
}

//js:get
func (e *AssertionErrorBase) Expected() any {
	return e.expected
}

//js:get
func (e *AssertionErrorBase) Operator() *string {
	return e.operator
}

//js:get
func (e *AssertionErrorBase) Code() string {
	return AssertionErrorCode
}

//js:get
func (e *AssertionErrorBase) GeneratedMessage() bool {
	return e.message == nil
}

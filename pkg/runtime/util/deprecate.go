//js:package util

package util

import (
	isolates "github.com/grexie/isolates"
)

//js:method deprecate
//js:export deprecate
func Deprecate(fn *isolates.Value, msg ...*isolates.Value) (isolates.Function, error) {
	return func(in isolates.FunctionArgs) (*isolates.Value, error) {
		isolates.For(in.ExecutionContext).Info("DEPRECATED")
		return fn.CallValue(in.ExecutionContext, in.This, in.Args...)
	}, nil
}

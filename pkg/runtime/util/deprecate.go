//js:package util

package util

//go:generate go run github.com/grexie/isolates/codegen

import (
	isolates "github.com/grexie/isolates"
)

//js:method deprecate
//js:export deprecate
//ts:export function deprecate<T extends Function>(fn: T, msg: string, code?: string): T;
func Deprecate(fn *isolates.Value, msg ...*isolates.Value) (isolates.Function, error) {
	return func(in isolates.FunctionArgs) (*isolates.Value, error) {
		isolates.For(in.ExecutionContext).Info("DEPRECATED")
		return fn.CallValue(in.ExecutionContext, in.This, in.Args...)
	}, nil
}

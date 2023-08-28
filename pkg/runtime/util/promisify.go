//js:package util

package util

import (
	"context"
	"fmt"

	isolates "github.com/grexie/isolates"
)

//js:method
//js:export promisify
func Promisify(ctx context.Context, fn *isolates.Value) (isolates.Function, error) {
	if !fn.IsKind(isolates.KindFunction) {
		return nil, fmt.Errorf("util.promisify: not a function %s", fn)
	}

	return func(in isolates.FunctionArgs) (*isolates.Value, error) {
		if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
			return nil, err
		} else if callback, err := in.Context.Create(in.ExecutionContext, func(in isolates.FunctionArgs) (*isolates.Value, error) {
			if in.Arg(in.ExecutionContext, 0).IsNil() {
				return nil, resolver.Resolve(in.ExecutionContext, in.Arg(in.ExecutionContext, 1))
			} else {
				return nil, resolver.Reject(in.ExecutionContext, in.Arg(in.ExecutionContext, 0))
			}
		}); err != nil {
			return nil, err
		} else {
			args := append(in.Args, callback)
			if _, err := fn.CallValue(in.ExecutionContext, in.This, args...); err != nil {
				return nil, err
			} else {
				return resolver.Promise(in.ExecutionContext)
			}
		}
	}, nil
}

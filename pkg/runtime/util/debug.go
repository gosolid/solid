//js:package util

package util

import (
	"context"

	"github.com/grexie/isolates"
)

//go:generate go run github.com/grexie/isolates/codegen

//js:export debuglog
//js:method debuglog
func Debuglog(ctx context.Context, section string, callback *isolates.Value) (*isolates.Value, error) {
	fn := func(in isolates.FunctionArgs) (*isolates.Value, error) {
		if global, err := in.Context.Global(in.ExecutionContext); err != nil {
			return nil, err
		} else if console, err := global.Get(in.ExecutionContext, "console"); err != nil {
			return nil, err
		} else if debug, err := console.Get(in.ExecutionContext, "debug"); err != nil {
			return nil, err
		} else if _, err := debug.CallValue(in.ExecutionContext, console, in.Args...); err != nil {
			return nil, err
		} else {
			return nil, nil
		}
	}

	if fnv, err := isolates.For(ctx).Context().Create(ctx, fn); err != nil {
		return nil, err
	} else {
		if callback.IsKind(isolates.KindFunction) {
			isolates.For(ctx).Context().AddMicrotask(ctx, func(in isolates.FunctionArgs) error {
				if _, err := callback.CallValue(in.ExecutionContext, nil, fnv); err != nil {
					return err
				} else {
					return nil
				}
			})
		}

		return fnv, nil
	}
}

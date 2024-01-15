//js:package util

package util

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"

	"github.com/grexie/isolates"
)

//js:method
//js:export inherits
func Inherits(ctx context.Context, constructor *isolates.Value, superConstructor *isolates.Value) (*isolates.Value, error) {
	if constructor.IsKind(isolates.KindFunction) && superConstructor.IsKind(isolates.KindFunction) {
		if err := constructor.Set(ctx, "super_", superConstructor); err != nil {
			return nil, err
		} else if superPrototype, err := superConstructor.Get(ctx, "prototype"); err != nil {
			return nil, err
		} else if prototype, err := isolates.For(ctx).Context().ObjectCreate(ctx, superPrototype, &map[string]*isolates.PropertyDescriptor{
			"constructor": {
				Value:        constructor,
				Enumerable:   false,
				Writable:     true,
				Configurable: true,
			},
		}); err != nil {
			return nil, err
		} else if err := constructor.Set(ctx, "prototype", prototype); err != nil {
			return nil, err
		}
	}

	return constructor, nil
}

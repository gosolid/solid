package solid

import (
	"context"

	"github.com/grexie/isolates"
)

type v8Value isolates.Value

func (v *v8Value) ToStringArray(ctx context.Context) ([]string, error) {
	if lengthV, err := (*isolates.Value)(v).Get(ctx, "length"); err != nil {
		return nil, err
	} else if length, err := lengthV.Int64(ctx); err != nil {
		return nil, err
	} else {
		arr := make([]string, length)

		for i := int64(0); i < length; i++ {
			if val, err := (*isolates.Value)(v).GetIndex(ctx, int(i)); err != nil {
				return nil, err
			} else if arr[i], err = val.StringValue(ctx); err != nil {
				return nil, err
			}
		}

		return arr, nil
	}
}

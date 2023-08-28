//js:package util

package util

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/grexie/isolates"
)

const kTab = "  "
const kDefaultOptions = "util.inspect.defaultOptions"

type Import struct{}

var formatRe = regexp.MustCompile("%(%|s|((\\d+)?(\\.(\\d+))?(d|i|f))|j|o|O|c)")

type InspectOptions struct {
	ShowHidden       bool `v8:"showHidden"`
	Depth            int  `v8:"depth"`
	Colors           bool `v8:"colors"`
	CustomInspect    bool `v8:"customInspect"`
	ShowProxy        bool `v8:"showProxy"`
	MaxArrayLength   int  `v8:"maxArrayLength"`
	MaxStringLength  int  `v8:"maxStringLength"`
	BreakLength      int  `v8:"breakLength"`
	Compact          any  `v8:"compact"`
	Sorted           any  `v8:"sorted"`
	Getters          any  `v8:"getters"`
	NumericSeparator bool `v8:"numericSeparator"`
}

//js:method
//js:export format
func Format(ctx context.Context, argsv ...any) (string, error) {
	return FormatWithOptions(ctx, nil, argsv...)
}

//js:method
//js:export formatWithOptions
func FormatWithOptions(ctx context.Context, options any, argsv ...any) (string, error) {
	if len(argsv) == 0 {
		return "", nil
	}

	format := argsv[0]
	argsv = argsv[1:]

	if formatv, ok := format.(*isolates.Value); ok && formatv.IsKind(isolates.KindString) {
		if f, err := formatv.StringValue(ctx); err != nil {
			return "", err
		} else {
			format = f
		}
	}

	if args, err := isolates.For(ctx).Context().CreateAll(ctx, argsv...); err != nil {
		return "", err
	} else if formats, ok := format.(string); ok {
		if len(args) == 0 {
			return formats, nil
		}

		i := 0
		msg := formatRe.ReplaceAllStringFunc(formats, func(s string) string {
			if s == "%%" {
				return "%"
			}

			var arg *isolates.Value
			if i < len(args) {
				arg = args[i]
				i++
			} else {
				return s
			}

			if s == "%s" {
				return fmt.Sprintf("%s", arg)
			} else if s == "%o" {
				if o, e := Inspect(ctx, arg, options); e != nil {
					err = e
					return ""
				} else {
					return o
				}
			} else if strings.HasSuffix(s, "d") || strings.HasSuffix(s, "i") {
				if arg.IsKind(isolates.KindNumber) {
					if d, err := arg.Int64(ctx); err != nil {
						return s
					} else {
						return fmt.Sprintf(s, d)
					}
				}
			} else if strings.HasSuffix(s, "f") {
				if arg.IsKind(isolates.KindNumber) {
					if f, err := arg.Float64(ctx); err != nil {
						return s
					} else {
						return fmt.Sprintf(s, f)
					}
				}
			}

			return s
		})

		if err != nil {
			return "", err
		}

		remainder := []string{msg}
		for _, arg := range args[i:] {
			if arg.IsKind(isolates.KindString) || arg.IsKind(isolates.KindStringObject) {
				if out, err := FormatWithOptions(ctx, options, "%s", arg); err != nil {
					return "", err
				} else {
					remainder = append(remainder, out)
				}
			} else if out, err := FormatWithOptions(ctx, options, "%o", arg); err != nil {
				return "", err
			} else {
				remainder = append(remainder, out)
			}
		}

		return strings.Join(remainder, " "), nil
	} else {
		argsv = append([]any{"%o", format}, argsv...)
		return FormatWithOptions(ctx, options, argsv...)
	}
}

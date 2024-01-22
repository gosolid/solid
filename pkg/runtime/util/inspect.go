//js:package util

package util

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/gosolid/solid/pkg/runtime/tty"
	"github.com/grexie/isolates"
)

//js:method inspect
//js:export inspect
func Inspect(ctx context.Context, object any, args ...any) (string, error) {
	var opts InspectOptions

	if defaultOptions, ok := isolates.For(ctx).Data(kDefaultOptions); !ok {
		if isolates.For(ctx).Context() == nil {
			return "", isolates.ErrNoContext
		} else {
			return "", fmt.Errorf("UNKNOWN ERROR")
		}
	} else if args, err := isolates.For(ctx).Context().CreateAll(ctx, args...); err != nil {
		return "", err
	} else {

		if len(args) == 1 && args[0].IsKind(isolates.KindObject) {
			if options, err := isolates.For(ctx).Context().Assign(ctx, map[string]any{}, defaultOptions, args[0]); err != nil {
				return "", err
			} else if options, err := options.Unmarshal(ctx, reflect.TypeOf(opts)); err != nil {
				return "", err
			} else {
				opts = options.Interface().(InspectOptions)
			}
		} else {
			if options, err := defaultOptions.(*isolates.Value).Unmarshal(ctx, reflect.TypeOf(opts)); err != nil {
				return "", err
			} else {
				opts = options.Interface().(InspectOptions)
			}

			if len(args) >= 1 && args[0].IsKind(isolates.KindBoolean) {
				if opts.ShowHidden, err = args[0].Bool(ctx); err != nil {
					return "", err
				}
			}

			if len(args) >= 2 && args[1].IsKind(isolates.KindBoolean) {
				if depth, err := args[1].Int64(ctx); err != nil {
					return "", err
				} else {
					opts.Depth = int(depth)
				}
			}

			if len(args) >= 3 && args[2].IsKind(isolates.KindBoolean) {
				if opts.Colors, err = args[2].Bool(ctx); err != nil {
					return "", err
				}
			}
		}
	}

	ansi := tty.ANSI

	if opts.Colors {
		ansi = ansi.Enable()
	} else {
		ansi = ansi.Disable()
	}

	context := serializeContext{
		options:      &opts,
		extraIndent:  "",
		depth:        0,
		seen:         map[any]bool{},
		literal:      ansi.Cyan().Format,
		key:          ansi.Black().Bold().Format,
		other:        ansi.White().Dim().Format,
		syntax:       ansi.White().Dim().Italic().Format,
		errorMessage: ansi.Red().Bold().Format,
		errorStack:   ansi.White().Dim().Format,
	}

	if objectv, err := isolates.For(ctx).Context().Create(ctx, object); err != nil {
		return "", err
	} else {
		return serialize(ctx, objectv, context)
	}
}

//js:method createInspectFunction
//js:export-instance inspect
func createInspectFunction(in isolates.RuntimeFunctionArgs) (*isolates.Value, error) {
	defaultOptions := &InspectOptions{
		Depth:           2,
		CustomInspect:   true,
		MaxArrayLength:  100,
		MaxStringLength: 10000,
		BreakLength:     80,
		Compact:         3,
		Sorted:          false,
		Getters:         false,
	}

	if defaultOptionsv, err := in.Context.Create(in.ExecutionContext, defaultOptions); err != nil {
		return nil, err
	} else if inspect, err := in.Exports.Get(in.ExecutionContext, "inspect"); err != nil {
		return nil, err
	} else if err := inspect.Set(in.ExecutionContext, "defaultOptions", defaultOptionsv); err != nil {
		return nil, err
	} else {
		in.Context.SetData(kDefaultOptions, defaultOptionsv)
		return inspect, nil
	}
}

type serializeContext struct {
	options     *InspectOptions
	depth       int
	extraIndent string
	seen        map[any]bool

	other        func(string) string
	syntax       func(string) string
	literal      func(string) string
	key          func(string) string
	errorMessage func(string) string
	errorStack   func(string) string
}

func (c *serializeContext) outerIndent() string {
	return strings.Repeat(kTab, c.depth) + c.extraIndent
}

func (c *serializeContext) innerIndent() string {
	return c.outerIndent() + kTab
}

func (c *serializeContext) lineBreak() string {
	return "\n"
}

func serialize(ctx context.Context, object *isolates.Value, context serializeContext) (string, error) {
	if object.IsKind(isolates.KindNull) {
		return context.literal("null"), nil
	} else if errorConstructor, err := isolates.For(ctx).Context().ErrorConstructor(ctx); err != nil {
		return "", err
	} else if object.InstanceOf(ctx, errorConstructor) {
		return serializeError(ctx, object, context)
	} else if object.IsKind(isolates.KindDate) {
		return serializeDate(ctx, object, context)
	} else if object.IsKind(isolates.KindFunction) {
		return serializeFunction(ctx, object, context)
	} else if object.IsKind(isolates.KindArray) {
		return serializeArray(ctx, object, context)
	} else if object.IsKind(isolates.KindObject) {
		return serializeObject(ctx, object, context)
	} else if object.IsKind(isolates.KindNumber) || object.IsKind(isolates.KindBoolean) || object.IsKind(isolates.KindString) || object.IsKind(isolates.KindStringObject) {
		return serializePrimitive(ctx, object, context)
	} else {
		return fmt.Sprintf("%s", object), nil
	}
}

func serializePrimitive(ctx context.Context, object *isolates.Value, context serializeContext) (string, error) {
	if object.IsKind(isolates.KindNumber) {
		if d, err := object.StringValue(ctx); err != nil {
			return "", err
		} else {
			return context.literal(fmt.Sprintf("%s", d)), nil
		}
	}

	if object.IsKind(isolates.KindBoolean) {
		if b, err := object.Bool(ctx); err != nil {
			return "", err
		} else if b {
			return context.literal("true"), nil
		} else {
			return context.literal("false"), nil
		}
	}

	if object.IsKind(isolates.KindString) || object.IsKind(isolates.KindStringObject) {
		if s, err := object.StringValue(ctx); err != nil {
			return "", err
		} else {
			s = strings.ReplaceAll(s, "\\", "\\\\")
			s = strings.ReplaceAll(s, "'", "\\'")
			return context.other("'") + context.literal(s) + context.other("'"), nil
		}
	}

	return "", fmt.Errorf("invalid primitive type")
}

func serializeDate(ctx context.Context, object *isolates.Value, context serializeContext) (string, error) {
	var t time.Time
	if rv, err := object.Unmarshal(ctx, reflect.TypeOf(t)); err != nil {
		return "", err
	} else {
		t = rv.Interface().(time.Time)
	}

	return context.syntax("date") + " " + context.literal(t.UTC().Format(time.RFC3339Nano)), nil
}

var reLineStart = regexp.MustCompile("^\\s+")

func serializeError(ctx context.Context, object *isolates.Value, context serializeContext) (string, error) {
	if stack, err := object.Get(ctx, "stack"); err != nil {
		return "", err
	} else if !stack.IsNil() && stack.IsKind(isolates.KindString) {
		var out strings.Builder

		lines := strings.Split(stack.String(), "\n")

		out.WriteString(context.errorMessage(lines[0]))

		for _, line := range lines[1:] {
			out.WriteString("\n")
			line = reLineStart.ReplaceAllLiteralString(line, context.innerIndent())
			out.WriteString(context.errorStack(line))
		}

		return out.String(), nil
	} else if message, err := object.Get(ctx, "message"); err != nil {
		return "", err
	} else if message.IsKind(isolates.KindString) {
		return context.errorMessage(fmt.Sprintf("%s", message)), nil
	} else {
		return context.errorMessage(fmt.Sprintf("%s", object)), nil
	}
}

func serializeFunction(ctx context.Context, object *isolates.Value, context serializeContext) (string, error) {
	if name, err := object.Get(ctx, "name"); err != nil {
		return "", err
	} else if name, err := name.StringValue(ctx); err != nil {
		return "", err
	} else {
		return context.syntax("function") + " " + context.literal(name) + context.other("(...)"), nil
	}
}

func serializeArray(ctx context.Context, object *isolates.Value, context serializeContext) (string, error) {
	if _, ok := context.seen[object]; ok {
		return context.other("[Circular]"), nil
	}
	context.seen[object] = true

	var out strings.Builder
	numLines := 0

	if context.depth < context.options.Depth {
		if length, err := object.GetLength(ctx); err != nil {
			return "", err
		} else {
			maxLength := int(math.Min(float64(length), float64(context.options.MaxArrayLength)))

			serialized := make([]string, maxLength)

			for i := 0; i < maxLength; i++ {
				childContext := context
				childContext.depth++

				if v, err := object.GetIndex(ctx, i); err != nil {
					return "", err
				} else if serialized[i], err = serialize(ctx, v, childContext); err != nil {
					return "", err
				}
			}

			var line strings.Builder

			for i := 0; i < len(serialized); i++ {
				lineLength := len(tty.ANSI.Strip(line.String() + serialized[i]))
				if line.Len() > 0 && lineLength >= context.options.BreakLength {
					if numLines > 0 {
						out.WriteString(context.lineBreak())
						out.WriteString(context.innerIndent())
					}
					numLines++
					out.WriteString(line.String())
					line.Reset()
				}
				line.WriteString(serialized[i])
				if i < len(serialized)-1 {
					line.WriteString(", ")
				}
			}
			if line.Len() > 0 {
				if numLines > 0 {
					out.WriteString(context.lineBreak())
					out.WriteString(context.innerIndent())
				}
				numLines++
				out.WriteString(line.String())
			}
		}
	}

	if context.depth >= context.options.Depth {
		return context.syntax("array"), nil
	} else if numLines <= 1 {
		return context.syntax("array") + " " + context.other("[") + out.String() + context.other("]"), nil
	} else {
		return context.syntax("array") + " " + context.other("[") + context.lineBreak() + context.innerIndent() + out.String() + context.lineBreak() + context.outerIndent() + context.other("]"), nil
	}
}

type descriptorInfo struct {
	isolates.PropertyDescriptor
	name         string
	constructors []*isolates.Value
	value        *isolates.Value
}

func serializeObject(ctx context.Context, object *isolates.Value, context serializeContext) (string, error) {
	descriptors := map[string]*descriptorInfo{}
	constructors := []*isolates.Value{}

	o := object
	for !o.IsNil() {
		if ownDescriptors, err := o.GetOwnPropertyDescriptors(ctx); err != nil {
			return "", err
		} else if constructor, err := o.Get(ctx, "constructor"); err != nil {
			return "", err
		} else {
			if o == object {
				if null, err := isolates.For(ctx).Context().Null(ctx); err != nil {
					return "", err
				} else {
					constructor = null
				}
			}

			for k, ownDescriptor := range ownDescriptors {
				if ownDescriptor.Enumerable || context.options.ShowHidden {
					if descriptor, ok := descriptors[k]; ok {
						hasConstructor := false

						for _, c := range descriptor.constructors {
							if equals, err := c.StrictEquals(ctx, constructor); err != nil {
								return "", err
							} else if equals {
								hasConstructor = true
								break
							}
						}

						if !hasConstructor {
							descriptor.constructors = append(descriptor.constructors, constructor)
						}
					} else if value, err := object.Get(ctx, k); err != nil {
						return "", err
					} else {
						descriptors[k] = &descriptorInfo{
							PropertyDescriptor: ownDescriptors[k],
							constructors:       []*isolates.Value{constructor},
							name:               k,
							value:              value,
						}
					}
				}
			}

			if !constructor.IsNil() {
				constructors = append(constructors, constructor)
			}
		}

		var err error
		if o, err = o.GetPrototype(ctx); err != nil {
			return "", err
		} else if o.IsNil() {
			break
		} else if nextPrototype, err := o.GetPrototype(ctx); err != nil {
			return "", err
		} else if nextPrototype.IsNil() {
			break
		}
	}

	entries := make([]*descriptorInfo, len(descriptors))
	i := 0
	for _, descriptor := range descriptors {
		entries[i] = descriptor
		i++
	}

	sort.Slice(entries, func(i, j int) bool {
		ci := -1
		cj := -1

		for m, constructor := range constructors {
			if len(entries[i].constructors) > 0 && constructor == entries[i].constructors[0] {
				ci = m
				break
			}
		}

		for m, constructor := range constructors {
			if len(entries[j].constructors) > 0 && constructor == entries[j].constructors[0] {
				cj = m
				break
			}
		}

		order := ci - cj
		if order == 0 && entries[i] != nil && entries[j] != nil {
			return sort.StringsAreSorted([]string{entries[i].name, entries[j].name})
		} else {
			return order < 0
		}
	})

	entries = append(entries, nil)

	var out strings.Builder

	out.WriteString(context.syntax("object"))

	if len(constructors) > 0 {
		if name, err := constructors[0].Get(ctx, "name"); err != nil || name.IsNil() {
			return "", err
		} else {
			out.WriteString(" ")
			out.WriteString(context.key(name.String()))
		}
	}

	if context.depth < context.options.Depth {
		out.WriteString(" ")
		out.WriteString(context.other("{"))
		if len(entries) > 1 {
			out.WriteString(context.lineBreak())
		}

		var previousConstructor *isolates.Value

		for _, descriptor := range entries {
			if descriptor == nil {
				if previousConstructor != nil {
					for _, c := range constructors {
						out.WriteString(context.innerIndent())
						out.WriteString(context.syntax("extends"))
						out.WriteString(" ")

						if name, err := c.Get(ctx, "name"); err != nil {
							return "", err
						} else if name.IsKind(isolates.KindString) {
							out.WriteString(context.other(name.String()))
						} else {
							out.WriteString(context.other("(anonymous)"))
						}

						if len(entries) > 1 {
							out.WriteString(context.lineBreak())
						}
					}
				}
				break
			}

			if descriptor.name == "constructor" {
				continue
			}

			if previousConstructor == nil || previousConstructor != descriptor.constructors[0] {
				previousConstructor = descriptor.constructors[0]

				if len(descriptor.constructors) > 0 && !descriptor.constructors[0].IsNil() {
					for i, c := range constructors {
						out.WriteString(context.innerIndent())
						out.WriteString(context.syntax("extends"))
						out.WriteString(" ")

						if name, err := c.Get(ctx, "name"); err != nil {
							return "", err
						} else if name.IsKind(isolates.KindString) {
							out.WriteString(context.other(name.String()))
						} else {
							out.WriteString(context.other("(anonymous)"))
						}

						out.WriteString(context.lineBreak())

						if descriptor.constructors[0] == c {
							constructors = constructors[i+1:]
							break
						}
					}
				}
			}

			overrides := func() error {
				var overrides []*isolates.Value
				for i, constructor := range descriptor.constructors {
					if equals, err := constructor.StrictEquals(ctx, previousConstructor); err != nil {
						return err
					} else if equals {
						overrides = descriptor.constructors[i+1:]
						break
					}
				}

				if overrides != nil && len(overrides) > 0 {
					out.WriteString(" ")
					out.WriteString(context.syntax("overrides"))
					out.WriteString(" ")
					for i, override := range overrides {
						if name, err := override.Get(ctx, "name"); err != nil {
							return err
						} else {
							if i != 0 {
								out.WriteString(context.other(", "))
							}
							out.WriteString(context.other(name.String()))
						}
					}
				}
				return nil
			}

			out.WriteString(context.innerIndent())
			if !previousConstructor.IsNil() {
				out.WriteString(kTab)
			}

			if descriptor.value.IsKind(isolates.KindFunction) {
				if name, err := descriptor.value.Get(ctx, "name"); err != nil {
					return "", err
				} else if name.String() == descriptor.name {
					out.WriteString(context.syntax("function"))
					out.WriteString(" ")
					out.WriteString(context.key(descriptor.name))
					out.WriteString(context.other("(...)"))
					if err := overrides(); err != nil {
						return "", err
					}
					out.WriteString(context.lineBreak())
					continue
				}
			}

			hasAccessors := false
			if !descriptor.Get.IsNil() {
				out.WriteString(context.syntax("get"))
				hasAccessors = true

				if !descriptor.Set.IsNil() {
					out.WriteString(context.syntax("/set"))
				}
			} else if !descriptor.Set.IsNil() {
				hasAccessors = true
				out.WriteString(context.syntax("set"))
			}

			if hasAccessors {
				out.WriteString(" ")
			}
			out.WriteString(context.key(descriptor.name))
			if hasAccessors {
				out.WriteString(context.other("()"))
			}

			out.WriteString(context.other(":"))
			out.WriteString(" ")

			childContext := context
			childContext.depth++
			if !previousConstructor.IsNil() {
				childContext.extraIndent += kTab
			}

			if serialized, err := serialize(ctx, descriptor.value, childContext); err != nil {
				return "", err
			} else {
				out.WriteString(serialized)
			}

			if err := overrides(); err != nil {
				return "", err
			}

			out.WriteString(context.lineBreak())
		}

		if len(entries) > 1 {
			out.WriteString(context.outerIndent())
		}
		out.WriteString(context.other("}"))
	}

	return out.String(), nil
}

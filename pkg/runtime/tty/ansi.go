//js:package tty

package tty

//go:generate go run github.com/grexie/isolates/codegen

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/grexie/isolates"
)

var ANSI = NewAnsiEscapeSequence()

type AnsiEscapeSequence struct {
	enabled bool
	flags   []int
}

//js:constructor ANSI
//js:export-instance ansi
func NewAnsiEscapeSequence() *AnsiEscapeSequence {
	return &AnsiEscapeSequence{
		enabled: true,
		flags:   []int{},
	}
}

//js:get enabled
func (e *AnsiEscapeSequence) Enabled() bool {
	return e.enabled
}

//js:method
func (e *AnsiEscapeSequence) Enable() *AnsiEscapeSequence {
	clone := *e
	clone.enabled = true
	return &clone
}

//js:method
func (e *AnsiEscapeSequence) Disable() *AnsiEscapeSequence {
	clone := *e
	clone.enabled = false
	return &clone
}

//js:method
func (e *AnsiEscapeSequence) Home() string {
	if !e.enabled {
		return ""
	}

	return "\033[H"
}

var seqRe = regexp.MustCompile("\033([a-zA-Z]|\\[\\d+(;\\d+)*[a-zA-Z]|\\[[a-zA-Z])")

//js:method
func (e *AnsiEscapeSequence) Strip(s string) string {
	return seqRe.ReplaceAllString(s, "")
}

//js:method
func (e *AnsiEscapeSequence) Move(line int, column int) string {
	if !e.enabled {
		return ""
	}

	return fmt.Sprintf("\033[%d;%dH", line, column)
}

//js:method
func (e *AnsiEscapeSequence) MoveUp(lines int) string {
	if !e.enabled {
		return ""
	}

	return fmt.Sprintf("\033[%dA", lines)
}

//js:method
func (e *AnsiEscapeSequence) MoveDown(lines int) string {
	if !e.enabled {
		return ""
	}

	return fmt.Sprintf("\033[%dB", lines)
}

//js:method
func (e *AnsiEscapeSequence) MoveRight(columns int) string {
	if !e.enabled {
		return ""
	}

	return fmt.Sprintf("\033[%dC", columns)
}

//js:method
func (e *AnsiEscapeSequence) MoveLeft(columns int) string {
	if !e.enabled {
		return ""
	}

	return fmt.Sprintf("\033[%dD", columns)
}

//js:method
func (e *AnsiEscapeSequence) MoveToStartDown(lines int) string {
	if !e.enabled {
		return ""
	}

	return fmt.Sprintf("\033[%dE", lines)
}

//js:method
func (e *AnsiEscapeSequence) MoveToStartUp(lines int) string {
	if !e.enabled {
		return ""
	}

	return fmt.Sprintf("\033[%dF", lines)
}

//js:method
func (e *AnsiEscapeSequence) MoveToStart() string {
	if !e.enabled {
		return ""
	}

	return "\033[0G"
}

//js:method
func (e *AnsiEscapeSequence) RequestCursorPosition() string {
	if !e.enabled {
		return ""
	}

	return "\033[6n"
}

//js:method
func (e *AnsiEscapeSequence) MoveUpOne() string {
	if !e.enabled {
		return ""
	}

	return "\033M"
}

//js:method
func (e *AnsiEscapeSequence) SaveCursorPosition() string {
	if !e.enabled {
		return ""
	}

	return "\0337"
}

//js:method
func (e *AnsiEscapeSequence) RestoreCursorPosition() string {
	if !e.enabled {
		return ""
	}

	return "\0338"
}

//js:method
func (e *AnsiEscapeSequence) SaveCursorPositionSCO() string {
	if !e.enabled {
		return ""
	}

	return "\033[s"
}

//js:method
func (e *AnsiEscapeSequence) RequestCursorPositionSCO() string {
	if !e.enabled {
		return ""
	}

	return "\033[u"
}

//js:method
func (e *AnsiEscapeSequence) EraseInDisplay() string {
	if !e.enabled {
		return ""
	}

	return "\033[J"
}

//js:method
func (e *AnsiEscapeSequence) EraseFromCursorToEndOfScreen() string {
	if !e.enabled {
		return ""
	}

	return "\033[0J"
}

//js:method
func (e *AnsiEscapeSequence) EraseFromCursorToStartOfScreen() string {
	if !e.enabled {
		return ""
	}

	return "\033[1J"
}

//js:method
func (e *AnsiEscapeSequence) EraseEntireScreen() string {
	if !e.enabled {
		return ""
	}

	return "\033[2J"
}

//js:method
func (e *AnsiEscapeSequence) EraseSavedLines() string {
	if !e.enabled {
		return ""
	}

	return "\033[3J"
}

//js:method
func (e *AnsiEscapeSequence) EraseInLine() string {
	if !e.enabled {
		return ""
	}

	return "\033[K"
}

//js:method
func (e *AnsiEscapeSequence) EraseFromCursorToEndOfLine() string {
	if !e.enabled {
		return ""
	}

	return "\033[0K"
}

//js:method
func (e *AnsiEscapeSequence) EraseFromCursorToStartOfLine() string {
	if !e.enabled {
		return ""
	}

	return "\033[1K"
}

//js:method
func (e *AnsiEscapeSequence) EraseLine() string {
	if !e.enabled {
		return ""
	}

	return "\033[2K"
}

func (e *AnsiEscapeSequence) append(flags ...int) *AnsiEscapeSequence {
	clone := *e
	clone.flags = append(e.flags, flags...)
	return &clone
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Reset() *AnsiEscapeSequence {
	return e.append(0)
}

var ansiReset = ANSI.Reset().String()

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Bold() *AnsiEscapeSequence {
	return e.append(1)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BoldReset() *AnsiEscapeSequence {
	return e.append(22)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Dim() *AnsiEscapeSequence {
	return e.append(2)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) DimReset() *AnsiEscapeSequence {
	return e.append(22)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Italic() *AnsiEscapeSequence {
	return e.append(3)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) ItalicReset() *AnsiEscapeSequence {
	return e.append(23)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Underline() *AnsiEscapeSequence {
	return e.append(4)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) UnderlineReset() *AnsiEscapeSequence {
	return e.append(24)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Blink() *AnsiEscapeSequence {
	return e.append(5)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BlinkReset() *AnsiEscapeSequence {
	return e.append(25)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Reverse() *AnsiEscapeSequence {
	return e.append(7)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) ReverseReset() *AnsiEscapeSequence {
	return e.append(27)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Hidden() *AnsiEscapeSequence {
	return e.append(8)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) HiddenReset() *AnsiEscapeSequence {
	return e.append(28)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Strikethrough() *AnsiEscapeSequence {
	return e.append(9)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) StrikethroughReset() *AnsiEscapeSequence {
	return e.append(29)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Black() *AnsiEscapeSequence {
	return e.append(30)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundBlack() *AnsiEscapeSequence {
	return e.append(40)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BrightBlack() *AnsiEscapeSequence {
	return e.append(90)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundBrightBlack() *AnsiEscapeSequence {
	return e.append(100)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Red() *AnsiEscapeSequence {
	return e.append(31)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundRed() *AnsiEscapeSequence {
	return e.append(41)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BrightRed() *AnsiEscapeSequence {
	return e.append(91)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundBrightRed() *AnsiEscapeSequence {
	return e.append(101)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Green() *AnsiEscapeSequence {
	return e.append(32)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundGreen() *AnsiEscapeSequence {
	return e.append(42)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BrightGreen() *AnsiEscapeSequence {
	return e.append(92)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundBrightGreen() *AnsiEscapeSequence {
	return e.append(102)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Yellow() *AnsiEscapeSequence {
	return e.append(33)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundYellow() *AnsiEscapeSequence {
	return e.append(43)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BrightYellow() *AnsiEscapeSequence {
	return e.append(93)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundBrightYellow() *AnsiEscapeSequence {
	return e.append(103)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Blue() *AnsiEscapeSequence {
	return e.append(34)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundBlue() *AnsiEscapeSequence {
	return e.append(44)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BrightBlue() *AnsiEscapeSequence {
	return e.append(94)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundBrightBlue() *AnsiEscapeSequence {
	return e.append(104)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Magenta() *AnsiEscapeSequence {
	return e.append(35)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundMagenta() *AnsiEscapeSequence {
	return e.append(45)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BrightMagenta() *AnsiEscapeSequence {
	return e.append(95)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundBrightMagenta() *AnsiEscapeSequence {
	return e.append(105)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Cyan() *AnsiEscapeSequence {
	return e.append(36)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundCyan() *AnsiEscapeSequence {
	return e.append(46)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BrightCyan() *AnsiEscapeSequence {
	return e.append(96)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundBrightCyan() *AnsiEscapeSequence {
	return e.append(106)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) White() *AnsiEscapeSequence {
	return e.append(37)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundWhite() *AnsiEscapeSequence {
	return e.append(47)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BrightWhite() *AnsiEscapeSequence {
	return e.append(97)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundBrightWhite() *AnsiEscapeSequence {
	return e.append(107)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) Default() *AnsiEscapeSequence {
	return e.append(39)
}

//js:get
//js:decorator getterDecorator
func (e *AnsiEscapeSequence) BackgroundDefault() *AnsiEscapeSequence {
	return e.append(49)
}

//js:method
//js:decorator functionDecorator
func (e *AnsiEscapeSequence) IndexedColor(index int) *AnsiEscapeSequence {
	return e.append(38, 5, index)
}

//js:method
//js:decorator functionDecorator
func (e *AnsiEscapeSequence) BackgroundIndexedColor(index int) *AnsiEscapeSequence {
	return e.append(48, 5, index)
}

//js:method
//js:decorator functionDecorator
func (e *AnsiEscapeSequence) Color(r int, g int, b int) *AnsiEscapeSequence {
	return e.append(38, 2, r, g, b)
}

//js:method
//js:decorator functionDecorator
func (e *AnsiEscapeSequence) BackgroundColor(r int, g int, b int) *AnsiEscapeSequence {
	return e.append(48, 2, r, g, b)
}

func (e *AnsiEscapeSequence) String() string {
	if !e.enabled || len(e.flags) == 0 {
		return ""
	}

	flags := make([]string, len(e.flags))

	for i, flag := range e.flags {
		flags[i] = fmt.Sprintf("%d", flag)
	}

	return fmt.Sprintf("\033[%sm", strings.Join(flags, ";"))
}

func (e *AnsiEscapeSequence) Format(s string) string {
	if !e.enabled {
		return s
	}

	reset := ANSI

	for _, flag := range e.flags {
		switch flag {
		case 1:
			reset = reset.BoldReset()
			break
		case 2:
			reset = reset.DimReset()
			break
		case 3:
			reset = reset.ItalicReset()
			break
		case 4:
			reset = reset.UnderlineReset()
			break
		case 5:
			reset = reset.BlinkReset()
			break
		case 7:
			reset = reset.ReverseReset()
			break
		case 8:
			reset = reset.HiddenReset()
			break
		case 9:
			reset = reset.StrikethroughReset()
			break
		}
	}
	seq := e.String()

	return fmt.Sprintf("%s%s%s%s", seq, strings.ReplaceAll(s, ansiReset, seq), reset.String(), ansiReset)
}

func (e *AnsiEscapeSequence) getterDecorator(in isolates.GetterArgs) (*isolates.Value, error) {
	if accessor, err := in.Context.Create(in.ExecutionContext, func(in isolates.FunctionArgs) (*isolates.Value, error) {
		if s, err := in.Arg(in.ExecutionContext, 0).StringValue(in.ExecutionContext); err != nil {
			return nil, err
		} else {
			return in.Context.Create(in.ExecutionContext, e.Format(s))
		}
	}); err != nil {
		return nil, err
	} else if object, err := in.Context.Create(in.ExecutionContext, e); err != nil {
		return nil, err
	} else if prototype, err := object.GetPrototype(in.ExecutionContext); err != nil {
		return nil, err
	} else if keys, err := object.Keys(in.ExecutionContext); err != nil {
		return nil, err
	} else if descriptors, err := prototype.GetOwnPropertyDescriptors(in.ExecutionContext); err != nil {
		return nil, err
	} else {
		accessor.SetReceiver(in.ExecutionContext, object.Receiver(in.ExecutionContext))

		for _, name := range keys {
			if getter, err := in.Context.Create(in.ExecutionContext, func(name string) isolates.Function {
				return func(in isolates.FunctionArgs) (*isolates.Value, error) {
					return object.Get(in.ExecutionContext, name)
				}
			}(name)); err != nil {
				return nil, err
			} else if undefined, err := in.Context.Undefined(in.ExecutionContext); err != nil {
				return nil, err
			} else {
				descriptors[name] = isolates.PropertyDescriptor{
					Get:          getter,
					Set:          undefined,
					Configurable: true,
				}
			}
		}

		for name, descriptor := range descriptors {
			if err := accessor.DefineProperty(in.ExecutionContext, name, &descriptor); err != nil {
				return nil, err
			}
		}

		return accessor, nil
	}
}

func (e *AnsiEscapeSequence) functionDecorator(in isolates.FunctionArgs) (*isolates.Value, error) {
	if accessor, err := in.Context.Create(in.ExecutionContext, func(in isolates.FunctionArgs) (*isolates.Value, error) {
		if s, err := in.Arg(in.ExecutionContext, 0).StringValue(in.ExecutionContext); err != nil {
			return nil, err
		} else {
			return in.Context.Create(in.ExecutionContext, e.Format(s))
		}
	}); err != nil {
		return nil, err
	} else if object, err := in.Context.Create(in.ExecutionContext, e); err != nil {
		return nil, err
	} else if prototype, err := object.GetPrototype(in.ExecutionContext); err != nil {
		return nil, err
	} else if keys, err := object.Keys(in.ExecutionContext); err != nil {
		return nil, err
	} else if descriptors, err := prototype.GetOwnPropertyDescriptors(in.ExecutionContext); err != nil {
		return nil, err
	} else {
		accessor.SetReceiver(in.ExecutionContext, object.Receiver(in.ExecutionContext))

		for _, name := range keys {
			if getter, err := in.Context.Create(in.ExecutionContext, func(name string) isolates.Function {
				return func(in isolates.FunctionArgs) (*isolates.Value, error) {
					return object.Get(in.ExecutionContext, name)
				}
			}(name)); err != nil {
				return nil, err
			} else if undefined, err := in.Context.Undefined(in.ExecutionContext); err != nil {
				return nil, err
			} else {
				descriptors[name] = isolates.PropertyDescriptor{
					Get:          getter,
					Set:          undefined,
					Configurable: true,
				}
			}
		}

		for name, descriptor := range descriptors {
			if err := accessor.DefineProperty(in.ExecutionContext, name, &descriptor); err != nil {
				return nil, err
			}
		}

		return accessor, nil
	}
}

//js:package console

package console

//go:generate go run github.com/grexie/isolates/codegen

import (
	"errors"
	"fmt"

	"github.com/gosolid/solid/pkg/runtime/util"
	"github.com/grexie/isolates"
)

var _ Console = &ConsoleBase{}

//js:alias ConsoleBase
type Console interface {
	Assert(args ...any)
	Clear()
	Count(args ...any)
	CountReset(args ...any)
	Debug(args ...any)
	Dir(args ...any)
	DirXML(args ...any)
	Error(args ...any)
	Group(args ...any)
	GroupCollapsed(args ...any)
	GroupEnd(args ...any)
	Info(args ...any)
	Log(args ...any)
	Table(args ...any)
	Time(args ...any)
	TimeEnd(args ...any)
	TimeLog(args ...any)
	Trace(args ...any)
	Warn(args ...any)
	Profile(args ...any)
	ProfileEnd(args ...any)
	TimeStamp(args ...any)
}

type ConsoleBase struct {
	context *isolates.Context
	stdout  *isolates.Value
	stderr  *isolates.Value
}

//js:constructor Console
//js:export Console
func NewConsole(in isolates.FunctionArgs) (*ConsoleBase, error) {
	var stdout, stderr *isolates.Value

	if len(in.Args) == 1 {
		var err error
		if stdout, err = in.Args[0].Get(in.ExecutionContext, "stdout"); err != nil {
			return nil, err
		} else if stderr, err = in.Args[0].Get(in.ExecutionContext, "stderr"); err != nil {
			return nil, err
		}
	} else if len(in.Args) >= 2 {
		stdout = in.Args[0]
		stderr = in.Args[1]
	} else {
		return nil, errors.New("invalid constructor args for Console")
	}

	in.This.RebindAll(in.ExecutionContext)

	return &ConsoleBase{
		context: in.Context,
		stdout:  stdout,
		stderr:  stderr,
	}, nil
}

//js:method createDefaultConsole
//js:export-instance default
func newDefaultConsole(in isolates.RuntimeFunctionArgs) (Console, error) {
	if taskModule, err := in.Require.Call(in.ExecutionContext, nil, "task"); err != nil {
		return nil, err
	} else if process, err := taskModule.Get(in.ExecutionContext, "process"); err != nil {
		return nil, err
	} else if stdout, err := process.Get(in.ExecutionContext, "stdout"); err != nil {
		return nil, err
	} else if stderr, err := process.Get(in.ExecutionContext, "stderr"); err != nil {
		return nil, err
	} else {
		if console, err := in.Context.New(in.ExecutionContext, NewConsole, stdout, stderr); err != nil {
			return nil, err
		} else {
			return console.(Console), nil
		}
	}
}

func (c *ConsoleBase) Write(writable *isolates.Value, msg string) {
	ctx := writable.GetContext().GetIsolate().GetExecutionContext()
	isolates.For(ctx).SetContext(c.context)
	writable.CallMethod(ctx, "write", msg)
}

func (c *ConsoleBase) WriteLine(writable *isolates.Value, msg string) {
	c.Write(writable, msg+"\n")
}

//js:method
func (c *ConsoleBase) Assert(args ...any) {

}

//js:method
func (c *ConsoleBase) Clear() {
	ctx := c.stderr.GetContext().GetIsolate().GetExecutionContext()
	isolates.For(ctx).SetContext(c.context)

	if isttyv, err := c.stderr.Get(ctx, "isTTY"); err != nil {
		return
	} else if istty, err := isttyv.Bool(ctx); err != nil {
		return
	} else if istty {
		c.stderr.CallMethod(ctx, "clear")
	}
}

//js:method
func (c *ConsoleBase) Count(args ...any) {

}

//js:method
func (c *ConsoleBase) CountReset(args ...any) {

}

//js:method
func (c *ConsoleBase) Debug(args ...any) {
	ctx := c.stdout.GetContext().GetIsolate().GetExecutionContext()
	isolates.For(ctx).SetContext(c.context)

	if msg, err := util.FormatWithOptions(ctx, map[string]any{"colors": true, "showHidden": true}, args...); err != nil {
		fmt.Println(err)
		return
	} else {
		c.WriteLine(c.stderr, msg)
	}
}

//js:method
func (c *ConsoleBase) Dir(args ...any) {

}

//js:method
func (c *ConsoleBase) DirXML(args ...any) {

}

//js:method
func (c *ConsoleBase) Error(args ...any) {
	ctx := c.stdout.GetContext().GetIsolate().GetExecutionContext()
	isolates.For(ctx).SetContext(c.context)

	if msg, err := util.FormatWithOptions(ctx, map[string]any{"colors": true, "showHidden": true}, args...); err != nil {
		fmt.Println(err)
		return
	} else {
		c.WriteLine(c.stderr, msg)
	}
}

//js:method
func (c *ConsoleBase) Group(args ...any) {

}

//js:method
func (c *ConsoleBase) GroupCollapsed(args ...any) {

}

//js:method
func (c *ConsoleBase) GroupEnd(args ...any) {

}

//js:method
func (c *ConsoleBase) Info(args ...any) {
	c.Log(args...)
}

//js:method
func (c *ConsoleBase) Log(args ...any) {
	ctx := c.context.GetIsolate().GetExecutionContext()
	isolates.For(ctx).SetContext(c.context)

	if msg, err := util.FormatWithOptions(ctx, map[string]any{"colors": true, "showHidden": true}, args...); err != nil {
		fmt.Println(err)
		return
	} else {
		c.WriteLine(c.stdout, msg)
	}
}

//js:method
func (c *ConsoleBase) Table(args ...any) {

}

//js:method
func (c *ConsoleBase) Time(args ...any) {

}

//js:method
func (c *ConsoleBase) TimeEnd(args ...any) {

}

//js:method
func (c *ConsoleBase) TimeLog(args ...any) {

}

//js:method
func (c *ConsoleBase) Trace(args ...any) {
	ctx := c.stdout.GetContext().GetIsolate().GetExecutionContext()
	isolates.For(ctx).SetContext(c.context)

	if msg, err := util.FormatWithOptions(ctx, map[string]any{"colors": true, "showHidden": true}, args...); err != nil {
		fmt.Println(err)
		return
	} else {
		c.WriteLine(c.stderr, msg)
	}
}

//js:method
func (c *ConsoleBase) Warn(args ...any) {
	c.Error(args...)
}

//js:method
func (c *ConsoleBase) Profile(args ...any) {

}

//js:method
func (c *ConsoleBase) ProfileEnd(args ...any) {

}

//js:method
func (c *ConsoleBase) TimeStamp(args ...any) {

}

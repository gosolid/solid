//js:package task
package solid

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/gosolid/solid/pkg/runtime"
	"github.com/gosolid/solid/pkg/runtime/events"
	"github.com/gosolid/solid/pkg/runtime/fs"
	"github.com/gosolid/solid/pkg/runtime/stream"
	"github.com/gosolid/solid/pkg/runtime/tty"
	"github.com/gosolid/solid/pkg/runtime/util"
	"github.com/grexie/isolates"
)

var _ runtime.Import

type Ticker struct {
	*time.Ticker
	StopC chan bool
}

//js:class Task
//js:export Task
type WorkerTask struct {
	startupMutex sync.Mutex
	mutex        sync.Mutex
	refCount     int64
	onTerminate  events.EventHandler
	exitCode     int

	options TaskOptions

	FS       fs.FS `v8:"fs"`
	path     string
	requires []string

	timeCounter *util.Counter
	timeEntered time.Time
	exitTimer   time.Timer

	stdout stream.Writable
	stderr stream.Writable
	stdin  stream.Readable

	isolate *isolates.Isolate
	context *isolates.Context
	handler *isolates.Value
	timer   *time.Timer

	env  map[string]string
	args []string

	nextTimer  int64
	timerMutex sync.Mutex
	timers     map[int64]*time.Timer

	nextTicker  int64
	tickerMutex sync.Mutex
	tickers     map[int64]*Ticker
}

//js:class Process
//js:export Process
type Process struct {
	*events.EventEmitterBase
	task *WorkerTask
}

type TaskOptions struct {
	FS        fs.FS
	Requires  []string
	Script    string
	Stdin     io.ReadCloser
	Stdout    io.WriteCloser
	Stderr    io.WriteCloser
	Env       map[string]string
	Arguments []string
}

func NewTask(options TaskOptions) (*WorkerTask, error) {

	task := &WorkerTask{
		options: options,

		FS:       options.FS,
		path:     options.Script,
		requires: options.Requires,

		timeCounter: util.NewCounter(60, 5*time.Minute),

		timers:  map[int64]*time.Timer{},
		tickers: map[int64]*Ticker{},

		env:  options.Env,
		args: options.Arguments,
	}

	return task, nil
}

func NewProcess(t *WorkerTask) *Process {
	return &Process{
		EventEmitterBase: &events.EventEmitterBase{},
		task:             t,
	}
}

func (p *Process) V8Construct(in isolates.FunctionArgs) {
	in.This.RebindAll(in.ExecutionContext)
}

func (t *WorkerTask) Clone() (*WorkerTask, error) {
	task := &WorkerTask{
		options:  t.options,
		FS:       t.FS,
		path:     t.path,
		requires: t.requires,

		timeCounter: util.NewCounter(60, 5*time.Minute),

		timers: map[int64]*time.Timer{},

		env: t.env,
	}

	return task, nil
}

func (t *WorkerTask) GetExecutionContext() context.Context {
	ctx := t.isolate.GetExecutionContext()
	isolates.For(ctx).SetContext(t.context)
	return ctx
}

//js:get exitCode
func (t *WorkerTask) ExitCode() int {
	return t.exitCode
}

//js:set exitCode
func (t *WorkerTask) SetExitCode(exitCode int) {
	t.exitCode = exitCode
}

//js:method
func (t *WorkerTask) Wait(ctx context.Context) (int, error) {
	if err := t.isolate.Wait(ctx); err != nil {
		if t.exitCode == 0 {
			return 1, err
		} else {
			return t.exitCode, err
		}
	} else {
		return t.exitCode, nil
	}
}

func (t *WorkerTask) OnTerminate(callback func() error) func() {
	return t.onTerminate.Add(func(...any) error {
		return callback()
	})
}

func (t *WorkerTask) Start() error {
	ctx := isolates.WithContext(context.Background())
	defer t.Terminate()

	if err := t.start(ctx); err != nil {
		return err
	} else {
		return nil
	}
}

func (t *WorkerTask) Stop() error {
	t.Terminate()
	return nil
}

func (c *WorkerTask) start(ctx context.Context) error {
	c.startupMutex.Lock()
	defer c.startupMutex.Unlock()

	if c.isolate != nil {
		return nil
	}

	var err error

	c.isolate = isolates.NewIsolate()

	c.isolate.AddExecutionEnterCallback(c.onEnter)
	c.isolate.AddExecutionExitCallback(c.onExit)

	if _, err := c.isolate.Sync(ctx, func(ctx context.Context) (interface{}, error) {
		c.isolate.AddShutdownHook(ctx, func(isolate *isolates.Isolate) {
			c.isolate = nil
			c.context = nil
			c.onTerminate.Emit()
		})

		if c.context, err = c.isolate.NewContext(ctx); err != nil {
			return nil, err
		} else if global, err := c.context.Global(ctx); err != nil {
			return nil, err
		} else if err := global.Set(ctx, "global", global); err != nil {
			return nil, err
		} else if _, err := c.context.RunWithRuntime(ctx, c.FS, c.path, func(in isolates.RuntimeFunctionArgs) error {
			if _, err := in.Require.Call(in.ExecutionContext, nil, "stream"); err != nil {
				return err
			}

			if stdout, err := c.context.New(in.ExecutionContext, tty.NewWriteStream, c.options.Stdout); err != nil {
				return err
			} else {
				c.stdout = stdout.(tty.WriteStream)
			}

			if stderr, err := c.context.New(in.ExecutionContext, tty.NewWriteStream, c.options.Stderr); err != nil {
				return err
			} else {
				c.stderr = stderr.(tty.WriteStream)
			}

			if stdin, err := c.context.New(in.ExecutionContext, tty.NewReadStream, c.options.Stdin); err != nil {
				return err
			} else {
				c.stdin = stdin.(tty.ReadStream)
			}

			process := NewProcess(c)

			if global, err := in.Context.Global(in.ExecutionContext); err != nil {
				return err
			} else if exports, err := in.Require.Call(in.ExecutionContext, nil, "task"); err != nil {
				return err
			} else if err := exports.Set(in.ExecutionContext, "task", c); err != nil {
				return err
			} else if err := exports.Set(in.ExecutionContext, "process", process); err != nil {
				return err
			} else if err := global.Set(in.ExecutionContext, "process", process); err != nil {
				return err
			}

			if _, err := in.Require.Call(in.ExecutionContext, nil, "util"); err != nil {
				return err
			} else if _, err := in.Require.Call(in.ExecutionContext, nil, "native:@grexie/workers/global"); err != nil {
				return err
			}

			for _, require := range c.requires {
				if _, err := in.Require.Call(in.ExecutionContext, nil, require); err != nil {
					return err
				}
			}

			return nil
		}, map[string]bool{}); err != nil {
			return nil, err
		} else {
			return nil, nil
		}
	}); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *WorkerTask) Ref() {
	c.refCount++
}

func (c *WorkerTask) Unref() {
	c.refCount--
	if c.refCount == 0 {
		c.Terminate()
	}
}

func (c *WorkerTask) onEnter() {
	c.timeEntered = time.Now()
}

func (c *WorkerTask) onExit() {
	c.timeCounter.AddDurationNow(c.timeEntered)
}

func (c *WorkerTask) Terminate() {
	if c.timer != nil {
		c.timer.Stop()
	}

	c.timer = time.AfterFunc(5000*time.Millisecond, func() {
		c.mutex.Lock()
		defer c.mutex.Unlock()

		if c.refCount == 0 {

			// c.isolate.Terminate()
		}
	})
}

func (c *WorkerTask) GetIsolate() *isolates.Isolate {
	return c.isolate
}

//js:method
func (c *WorkerTask) SetTimeout(callback *isolates.Value, duration time.Duration) (int64, error) {
	c.timerMutex.Lock()
	defer c.timerMutex.Unlock()

	timerId := c.nextTimer
	c.nextTimer++
	c.timers[timerId] = time.AfterFunc(duration, func() {
		ctx := isolates.WithContext(context.Background())
		isolates.For(ctx).SetContext(c.context)

		c.context.AddMicrotask(ctx, func(in isolates.FunctionArgs) error {
			c.timerMutex.Lock()
			if _, ok := c.timers[timerId]; !ok {
				c.timerMutex.Unlock()
				return nil
			}

			delete(c.timers, timerId)
			c.timerMutex.Unlock()

			if _, err := callback.CallValue(in.ExecutionContext, nil); err != nil {
				return err
			}

			return nil
		})

	})

	return timerId, nil
}

//js:method
func (c *WorkerTask) ClearTimeout(timer int64) {
	if t, ok := c.timers[timer]; ok {
		delete(c.timers, timer)
		t.Stop()
	}
}

//js:method
func (c *WorkerTask) SetInterval(ctx context.Context, callback *isolates.Value, duration time.Duration) (int64, error) {
	c.tickerMutex.Lock()
	defer c.tickerMutex.Unlock()

	tickerId := c.nextTicker
	c.nextTicker++
	ticker := &Ticker{
		Ticker: time.NewTicker(duration),
		StopC:  make(chan bool),
	}
	c.tickers[tickerId] = ticker

	isolates.For(ctx).Background(func(ctx context.Context) {
		for {
			select {
			case <-ticker.C:
				callback.Call(ctx, nil)
				break
			case <-ticker.StopC:
				return
			}
		}
	})

	return tickerId, nil
}

//js:method
func (c *WorkerTask) ClearInterval(ticker int64) {
	if t, ok := c.tickers[ticker]; ok {
		delete(c.tickers, ticker)
		t.Stop()
		t.StopC <- true
	}
}

func (c *WorkerTask) V8FuncSetImmediate(in isolates.FunctionArgs) (*isolates.Value, error) {
	fn := in.Arg(in.ExecutionContext, 0)
	var args []*isolates.Value
	if len(in.Args) > 1 {
		args = in.Args[1:]
	}

	in.Context.AddMicrotask(in.ExecutionContext, func(in isolates.FunctionArgs) error {
		if _, err := fn.CallValue(in.ExecutionContext, nil, args...); err != nil {
			return err
		}

		return nil
	})

	return nil, nil
}

func (c *WorkerTask) V8FuncCreateTask(taskFS fs.FS, requires []string, taskPath string) (*WorkerTask, error) {
	options := c.options
	options.Requires = requires
	options.FS = taskFS
	options.Script = taskPath

	return NewTask(options)
}

//js:method cwd
func (c *Process) Cwd() string {
	return "/app"
}

//js:get env
func (c *Process) Env() map[string]string {
	return c.task.env
}

//js:get args
func (c *Process) Args() []string {
	return c.task.args
}

//js:get platform
func (c *Process) Platform() string {
	return "workers"
}

type Versions struct {
	Node string `v8:"node"`
}

//js:get versions
func (c *Process) Versions() *Versions {
	return &Versions{
		Node: "18.0.0",
	}
}

//js:get
func (c *Process) Task() *WorkerTask {
	return c.task
}

//js:get stdout
func (c *Process) Stdout() stream.Writable {
	return c.task.stdout
}

//js:get stderr
func (c *Process) Stderr() stream.Writable {
	return c.task.stderr
}

//js:get argv
func (c *Process) Argv() []string {
	return c.task.args
}

//js:get stdin
func (c *Process) Stdin() stream.Readable {
	return c.task.stdin
}

//js:method
func (c *Process) NextTick(in isolates.FunctionArgs) (*isolates.Value, error) {
	fn := in.Arg(in.ExecutionContext, 0)
	args := in.Args[1:]

	in.Context.AddMicrotask(in.ExecutionContext, func(in isolates.FunctionArgs) error {
		if _, err := fn.CallValue(in.ExecutionContext, nil, args...); err != nil {
			return err
		}

		return nil
	})

	return nil, nil
}

//js:method
func (c *Process) Exit(exitCode int) error {
	c.task.exitCode = exitCode
	return c.task.Stop()
}

type WorkerTaskStats struct {
	Contexts                int           `v8:"contexts"`
	Values                  int           `v8:"values"`
	Receivers               int           `v8:"receivers"`
	ComputeDuration         time.Duration `v8:"computeDuration"`
	ComputeDuration5        time.Duration `v8:"computeDuration5"`
	ComputeDuration15       time.Duration `v8:"computeDuration15"`
	ComputeDuration30       time.Duration `v8:"computeDuration30"`
	ComputeDuration60       time.Duration `v8:"computeDuration60"`
	ComputeDuration120      time.Duration `v8:"computeDuration120"`
	DoesZapGarbage          bool          `v8:"doesZapGarbage"`
	HeapSizeLimit           uint64        `v8:"heapSizeLimit"`
	MallocedMemory          uint64        `v8:"mallocedMemory"`
	PeakMallocedMemory      uint64        `v8:"peakMallocedMemory"`
	TotalAvailableSize      uint64        `v8:"totalAvailableSize"`
	TotalHeapSize           uint64        `v8:"totalHeapSize"`
	TotalHeapSizeExecutable uint64        `v8:"totalHeapSizeExecutable"`
	TotalPhysicalSize       uint64        `v8:"totalPhysicalSize"`
	UsedHeapSize            uint64        `v8:"usedHeapSize"`
}

//js:constructor TaskStats
//js:export TaskStats
func NewWorkerTaskStats() *WorkerTaskStats {
	return &WorkerTaskStats{}
}

//js:get
func (c *WorkerTask) TimeCounter(ctx context.Context) (*util.Counter, error) {
	return c.timeCounter.Clone(ctx)
}

//js:get
func (c *WorkerTask) Stats(ctx context.Context) (*isolates.Value, error) {
	out := &WorkerTaskStats{}

	contexts, _ := c.isolate.Contexts(ctx)

	out.Contexts = len(contexts)
	for _, context := range contexts {
		out.Values += context.Values()
		out.Receivers += context.Receivers()
	}

	stats, _ := c.isolate.GetHeapStatistics(ctx)
	out.DoesZapGarbage = stats.DoesZapGarbage
	out.HeapSizeLimit = (stats.HeapSizeLimit)
	out.MallocedMemory = (stats.MallocedMemory)
	out.PeakMallocedMemory = (stats.PeakMallocedMemory)
	out.TotalAvailableSize = (stats.TotalAvailableSize)
	out.TotalHeapSize = (stats.TotalHeapSize)
	out.TotalHeapSizeExecutable = (stats.TotalHeapSizeExecutable)
	out.TotalPhysicalSize = (stats.TotalPhysicalSize)
	out.UsedHeapSize = (stats.UsedHeapSize)

	out.ComputeDuration = c.timeCounter.SumAllDurationNow()
	out.ComputeDuration5 = c.timeCounter.SumDurationNow(5 * time.Second)
	out.ComputeDuration15 = c.timeCounter.SumDurationNow(15 * time.Second)
	out.ComputeDuration30 = c.timeCounter.SumDurationNow(30 * time.Second)
	out.ComputeDuration60 = c.timeCounter.SumDurationNow(60 * time.Second)
	out.ComputeDuration120 = c.timeCounter.SumDurationNow(120 * time.Second)

	return isolates.For(ctx).Context().Create(ctx, out)
}

//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSThread : objc.NSObject
*/

type NSThread struct {
  *objc.NSObject

}

func (r *NSThread) CallStackSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) MainThread() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) CallStackReturnAddresses() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) IsMainThread() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) IsCancelled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) InitWithTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) Cancel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) CurrentThread() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) StackSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) IsExecuting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) ThreadPriority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) Start() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) QualityOfService() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) SetQualityOfService() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) DetachNewThreadWithBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) SleepForTimeInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) ThreadDictionary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) SetStackSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) IsFinished() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) DetachNewThreadSelector() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) SleepUntilDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) SetThreadPriority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) InitWithBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) Main() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) IsMultiThreaded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) Exit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


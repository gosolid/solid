//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSThread : objc.NSObject
*/

type NSThread struct {
  *objc.NSObject

}

func (r *NSThread) CurrentThread() (*NSThread, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSThread) MainThread() (*NSThread, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) DetachNewThreadSelector() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSThread) ThreadPriority() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSThread) CallStackReturnAddresses() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) Exit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSThread) Cancel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSThread) Start() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSThread) StackSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSThread) InitWithTarget() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) IsExecuting() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSThread) IsCancelled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSThread) SetStackSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSThread) SetThreadPriority() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSThread) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) InitWithBlock() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) Main() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSThread) SetQualityOfService() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSThread) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) IsMainThread() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSThread) SleepUntilDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSThread) IsFinished() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSThread) QualityOfService() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSThread) ThreadDictionary() (*NSMutableDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) DetachNewThreadWithBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSThread) SleepForTimeInterval() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSThread) CallStackSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSThread) IsMultiThreaded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


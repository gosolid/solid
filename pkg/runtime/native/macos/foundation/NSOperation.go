//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSOperation : objc.NSObject
*/

type NSOperation struct {
  *objc.NSObject

}

func (r *NSOperation) Cancel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperation) SetCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperation) SetThreadPriority() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperation) AddDependency() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperation) IsExecuting() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOperation) QueuePriority() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOperation) IsFinished() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOperation) IsAsynchronous() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOperation) CompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) ThreadPriority() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOperation) Start() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperation) Main() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperation) RemoveDependency() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperation) WaitUntilFinished() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperation) SetQueuePriority() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperation) QualityOfService() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOperation) SetQualityOfService() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperation) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) IsCancelled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOperation) IsConcurrent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOperation) IsReady() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOperation) Dependencies() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) SetName() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSOperation : objc.NSObject
*/

type NSOperation struct {
  *objc.NSObject

}

func (r *NSOperation) IsAsynchronous() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) IsReady() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) Dependencies() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) SetQueuePriority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) CompletionBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) Start() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) Main() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) IsConcurrent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) QualityOfService() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) SetQualityOfService() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) Cancel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) IsExecuting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) SetCompletionBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) ThreadPriority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) SetThreadPriority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) WaitUntilFinished() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) IsCancelled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) QueuePriority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) AddDependency() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) RemoveDependency() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperation) IsFinished() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


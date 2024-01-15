//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSOperationQueue : objc.NSObject
*/

type NSOperationQueue struct {
  *objc.NSObject

}

func (r *NSOperationQueue) SetMaxConcurrentOperationCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) UnderlyingQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) AddOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) AddOperations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) AddBarrierBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) AddOperationWithBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) MaxConcurrentOperationCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) MainQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) IsSuspended() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) SetSuspended() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) QualityOfService() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) SetQualityOfService() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) SetUnderlyingQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) CancelAllOperations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) WaitUntilAllOperationsAreFinished() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) Progress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) CurrentQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


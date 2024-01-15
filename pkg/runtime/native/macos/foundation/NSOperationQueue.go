//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSOperationQueue : objc.NSObject
*/

type NSOperationQueue struct {
  *objc.NSObject
  *NSProgressReporting
}

func (r *NSOperationQueue) AddOperations() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) AddBarrierBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) WaitUntilAllOperationsAreFinished() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) Progress() (*NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) IsSuspended() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) SetUnderlyingQueue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) AddOperationWithBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) SetSuspended() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) SetQualityOfService() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) AddOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) CancelAllOperations() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) MaxConcurrentOperationCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) CurrentQueue() (*NSOperationQueue, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) MainQueue() (*NSOperationQueue, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) SetMaxConcurrentOperationCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) QualityOfService() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOperationQueue) UnderlyingQueue() (**objc.NSObject, error) {
  return nil, fmt.Errorf("unimplemented")
}


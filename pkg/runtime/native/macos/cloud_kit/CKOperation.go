//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKOperation : Foundation.NSOperation
*/

type CKOperation struct {
  *foundation.NSOperation

}

func (r *CKOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKOperation) Configuration() (*CKOperationConfiguration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKOperation) SetConfiguration() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKOperation) Group() (*CKOperationGroup, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKOperation) SetGroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKOperation) OperationID() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKOperation) LongLivedOperationWasPersistedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKOperation) SetLongLivedOperationWasPersistedBlock() error {
  return fmt.Errorf("unimplemented")
}


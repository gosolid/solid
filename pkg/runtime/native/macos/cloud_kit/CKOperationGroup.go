//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKOperationGroup : objc.NSObject
*/

type CKOperationGroup struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *CKOperationGroup) DefaultConfiguration() (*CKOperationConfiguration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKOperationGroup) SetExpectedReceiveSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKOperationGroup) SetDefaultConfiguration() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKOperationGroup) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKOperationGroup) SetQuantity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKOperationGroup) OperationGroupID() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKOperationGroup) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKOperationGroup) ExpectedSendSize() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKOperationGroup) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKOperationGroup) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKOperationGroup) Quantity() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKOperationGroup) SetExpectedSendSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKOperationGroup) ExpectedReceiveSize() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


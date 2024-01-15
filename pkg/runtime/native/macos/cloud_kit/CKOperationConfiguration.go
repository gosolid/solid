//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CloudKit.CKOperationConfiguration : objc.NSObject
*/

type CKOperationConfiguration struct {
  *objc.NSObject

}

func (r *CKOperationConfiguration) SetQualityOfService() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKOperationConfiguration) AllowsCellularAccess() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKOperationConfiguration) SetAllowsCellularAccess() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKOperationConfiguration) SetLongLived() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKOperationConfiguration) QualityOfService() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKOperationConfiguration) SetContainer() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKOperationConfiguration) IsLongLived() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKOperationConfiguration) TimeoutIntervalForRequest() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKOperationConfiguration) SetTimeoutIntervalForRequest() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKOperationConfiguration) TimeoutIntervalForResource() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKOperationConfiguration) SetTimeoutIntervalForResource() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKOperationConfiguration) Container() (*CKContainer, error) {
  return nil, fmt.Errorf("unimplemented")
}


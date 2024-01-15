//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CloudKit.CKSystemSharingUIObserver : objc.NSObject
*/

type CKSystemSharingUIObserver struct {
  *objc.NSObject

}

func (r *CKSystemSharingUIObserver) SetSystemSharingUIDidStopSharingBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSystemSharingUIObserver) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSystemSharingUIObserver) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSystemSharingUIObserver) InitWithContainer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSystemSharingUIObserver) SystemSharingUIDidSaveShareBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSystemSharingUIObserver) SetSystemSharingUIDidSaveShareBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSystemSharingUIObserver) SystemSharingUIDidStopSharingBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}


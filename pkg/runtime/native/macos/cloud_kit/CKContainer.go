//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKContainer : objc.NSObject
*/

type CKContainer struct {
  *objc.NSObject

}

func (r *CKContainer) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKContainer) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKContainer) DefaultContainer() (*CKContainer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKContainer) ContainerWithIdentifier() (*CKContainer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKContainer) AddOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKContainer) ContainerIdentifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}


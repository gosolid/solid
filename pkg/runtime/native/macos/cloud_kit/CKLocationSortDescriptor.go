//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_location"
)

/*
interface CloudKit.CKLocationSortDescriptor : Foundation.NSSortDescriptor
*/

type CKLocationSortDescriptor struct {
  *foundation.NSSortDescriptor
  *foundation.NSSecureCoding
}

func (r *CKLocationSortDescriptor) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKLocationSortDescriptor) RelativeLocation() (*core_location.CLLocation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKLocationSortDescriptor) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKLocationSortDescriptor) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKLocationSortDescriptor) InitWithKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKQuery : objc.NSObject
*/

type CKQuery struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKQuery) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQuery) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQuery) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQuery) InitWithRecordType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQuery) RecordType() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQuery) Predicate() (*foundation.NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQuery) SortDescriptors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQuery) SetSortDescriptors() error {
  return fmt.Errorf("unimplemented")
}


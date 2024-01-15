//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKReference : objc.NSObject
*/

type CKReference struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKReference) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKReference) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKReference) InitWithRecordID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKReference) InitWithRecord() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKReference) ReferenceAction() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKReference) RecordID() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}


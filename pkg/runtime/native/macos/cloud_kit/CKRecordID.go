//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CloudKit.CKRecordID : objc.NSObject
*/

type CKRecordID struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKRecordID) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordID) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordID) InitWithRecordName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordID) RecordName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordID) ZoneID() (*CKRecordZoneID, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKRecordZoneID : objc.NSObject
*/

type CKRecordZoneID struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKRecordZoneID) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZoneID) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZoneID) InitWithZoneName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZoneID) ZoneName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZoneID) OwnerName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}


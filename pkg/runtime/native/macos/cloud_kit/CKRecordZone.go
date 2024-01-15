//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CloudKit.CKRecordZone : objc.NSObject
*/

type CKRecordZone struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKRecordZone) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZone) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZone) InitWithZoneName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZone) InitWithZoneID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZone) ZoneID() (*CKRecordZoneID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZone) Capabilities() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKRecordZone) Share() (*CKReference, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZone) DefaultRecordZone() (*CKRecordZone, error) {
  return nil, fmt.Errorf("unimplemented")
}


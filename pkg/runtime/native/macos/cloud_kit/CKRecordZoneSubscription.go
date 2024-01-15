//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKRecordZoneSubscription : CloudKit.CKSubscription
*/

type CKRecordZoneSubscription struct {
  *CKSubscription
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKRecordZoneSubscription) InitWithZoneID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZoneSubscription) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZoneSubscription) ZoneID() (*CKRecordZoneID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZoneSubscription) RecordType() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZoneSubscription) SetRecordType() error {
  return fmt.Errorf("unimplemented")
}


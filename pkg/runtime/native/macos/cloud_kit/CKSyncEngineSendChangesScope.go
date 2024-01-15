//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKSyncEngineSendChangesScope : objc.NSObject
*/

type CKSyncEngineSendChangesScope struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *CKSyncEngineSendChangesScope) InitWithRecordIDs() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSendChangesScope) ContainsRecordID() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSendChangesScope) ContainsPendingRecordZoneChange() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSendChangesScope) ZoneIDs() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSendChangesScope) ExcludedZoneIDs() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSendChangesScope) RecordIDs() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSendChangesScope) InitWithZoneIDs() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSendChangesScope) InitWithExcludedZoneIDs() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


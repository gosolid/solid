//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKSyncEngineSentRecordZoneChangesEvent : CloudKit.CKSyncEngineEvent
*/

type CKSyncEngineSentRecordZoneChangesEvent struct {
  *CKSyncEngineEvent

}

func (r *CKSyncEngineSentRecordZoneChangesEvent) DeletedRecordIDs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSentRecordZoneChangesEvent) FailedRecordDeletes() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSentRecordZoneChangesEvent) SavedRecords() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSentRecordZoneChangesEvent) FailedRecordSaves() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


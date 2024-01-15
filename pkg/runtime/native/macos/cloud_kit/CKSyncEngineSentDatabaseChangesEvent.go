//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKSyncEngineSentDatabaseChangesEvent : CloudKit.CKSyncEngineEvent
*/

type CKSyncEngineSentDatabaseChangesEvent struct {
  *CKSyncEngineEvent

}

func (r *CKSyncEngineSentDatabaseChangesEvent) SavedZones() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSentDatabaseChangesEvent) FailedZoneSaves() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSentDatabaseChangesEvent) DeletedZoneIDs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSentDatabaseChangesEvent) FailedZoneDeletes() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}


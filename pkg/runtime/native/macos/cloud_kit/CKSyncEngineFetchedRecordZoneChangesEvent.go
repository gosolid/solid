//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKSyncEngineFetchedRecordZoneChangesEvent : CloudKit.CKSyncEngineEvent
*/

type CKSyncEngineFetchedRecordZoneChangesEvent struct {
  *CKSyncEngineEvent

}

func (r *CKSyncEngineFetchedRecordZoneChangesEvent) Deletions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchedRecordZoneChangesEvent) Modifications() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


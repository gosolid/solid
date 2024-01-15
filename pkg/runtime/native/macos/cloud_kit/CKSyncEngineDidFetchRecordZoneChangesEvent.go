//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKSyncEngineDidFetchRecordZoneChangesEvent : CloudKit.CKSyncEngineEvent
*/

type CKSyncEngineDidFetchRecordZoneChangesEvent struct {
  *CKSyncEngineEvent

}

func (r *CKSyncEngineDidFetchRecordZoneChangesEvent) ZoneID() (*CKRecordZoneID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineDidFetchRecordZoneChangesEvent) Error() (*foundation.NSError, error) {
  return nil, fmt.Errorf("unimplemented")
}


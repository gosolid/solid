//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CloudKit.CKSyncEngineWillFetchRecordZoneChangesEvent : CloudKit.CKSyncEngineEvent
*/

type CKSyncEngineWillFetchRecordZoneChangesEvent struct {
  *CKSyncEngineEvent

}

func (r *CKSyncEngineWillFetchRecordZoneChangesEvent) ZoneID() (*CKRecordZoneID, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CloudKit.CKSyncEnginePendingZoneSave : CloudKit.CKSyncEnginePendingDatabaseChange
*/

type CKSyncEnginePendingZoneSave struct {
  *CKSyncEnginePendingDatabaseChange

}

func (r *CKSyncEnginePendingZoneSave) InitWithZone() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEnginePendingZoneSave) Zone() (*CKRecordZone, error) {
  return nil, fmt.Errorf("unimplemented")
}


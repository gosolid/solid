//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CloudKit.CKSyncEnginePendingZoneDelete : CloudKit.CKSyncEnginePendingDatabaseChange
*/

type CKSyncEnginePendingZoneDelete struct {
  *CKSyncEnginePendingDatabaseChange

}

func (r *CKSyncEnginePendingZoneDelete) InitWithZoneID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


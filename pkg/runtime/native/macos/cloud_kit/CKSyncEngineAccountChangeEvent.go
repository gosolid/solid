//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CloudKit.CKSyncEngineAccountChangeEvent : CloudKit.CKSyncEngineEvent
*/

type CKSyncEngineAccountChangeEvent struct {
  *CKSyncEngineEvent

}

func (r *CKSyncEngineAccountChangeEvent) ChangeType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineAccountChangeEvent) PreviousUser() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineAccountChangeEvent) CurrentUser() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}


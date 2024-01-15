//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CloudKit.CKSyncEnginePendingRecordZoneChange : objc.NSObject
*/

type CKSyncEnginePendingRecordZoneChange struct {
  *objc.NSObject

}

func (r *CKSyncEnginePendingRecordZoneChange) InitWithRecordID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEnginePendingRecordZoneChange) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEnginePendingRecordZoneChange) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEnginePendingRecordZoneChange) RecordID() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEnginePendingRecordZoneChange) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


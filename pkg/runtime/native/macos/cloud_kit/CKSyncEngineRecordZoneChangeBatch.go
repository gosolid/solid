//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKSyncEngineRecordZoneChangeBatch : objc.NSObject
*/

type CKSyncEngineRecordZoneChangeBatch struct {
  *objc.NSObject

}

func (r *CKSyncEngineRecordZoneChangeBatch) RecordsToSave() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineRecordZoneChangeBatch) RecordIDsToDelete() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineRecordZoneChangeBatch) AtomicByZone() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineRecordZoneChangeBatch) SetAtomicByZone() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineRecordZoneChangeBatch) InitWithPendingChanges() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineRecordZoneChangeBatch) InitWithRecordsToSave() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineRecordZoneChangeBatch) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineRecordZoneChangeBatch) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKSyncEngineState : objc.NSObject
*/

type CKSyncEngineState struct {
  *objc.NSObject

}

func (r *CKSyncEngineState) PendingDatabaseChanges() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineState) HasPendingUntrackedChanges() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineState) SetHasPendingUntrackedChanges() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineState) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineState) AddPendingRecordZoneChanges() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineState) AddPendingDatabaseChanges() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineState) RemovePendingDatabaseChanges() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineState) PendingRecordZoneChanges() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineState) ZoneIDsWithUnfetchedServerChanges() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineState) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineState) RemovePendingRecordZoneChanges() error {
  return fmt.Errorf("unimplemented")
}


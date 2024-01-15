//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKModifyRecordZonesOperation : CloudKit.CKDatabaseOperation
*/

type CKModifyRecordZonesOperation struct {
  *CKDatabaseOperation

}

func (r *CKModifyRecordZonesOperation) SetRecordZonesToSave() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordZonesOperation) RecordZoneIDsToDelete() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordZonesOperation) SetPerRecordZoneSaveBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordZonesOperation) SetPerRecordZoneDeleteBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordZonesOperation) ModifyRecordZonesCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordZonesOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordZonesOperation) InitWithRecordZonesToSave() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordZonesOperation) RecordZonesToSave() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordZonesOperation) SetModifyRecordZonesCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordZonesOperation) SetRecordZoneIDsToDelete() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordZonesOperation) PerRecordZoneSaveBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordZonesOperation) PerRecordZoneDeleteBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}


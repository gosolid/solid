//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKModifyRecordsOperation : CloudKit.CKDatabaseOperation
*/

type CKModifyRecordsOperation struct {
  *CKDatabaseOperation

}

func (r *CKModifyRecordsOperation) SetRecordsToSave() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) SavePolicy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) SetSavePolicy() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) PerRecordSaveBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) RecordsToSave() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) RecordIDsToDelete() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) SetAtomic() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) PerRecordCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) SetPerRecordCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) SetPerRecordSaveBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) InitWithRecordsToSave() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) Atomic() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) SetPerRecordProgressBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) PerRecordDeleteBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) SetPerRecordDeleteBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) SetModifyRecordsCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) SetRecordIDsToDelete() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) ClientChangeTokenData() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) SetClientChangeTokenData() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) PerRecordProgressBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyRecordsOperation) ModifyRecordsCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}


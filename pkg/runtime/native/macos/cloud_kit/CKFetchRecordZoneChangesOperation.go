//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKFetchRecordZoneChangesOperation : CloudKit.CKDatabaseOperation
*/

type CKFetchRecordZoneChangesOperation struct {
  *CKDatabaseOperation

}

func (r *CKFetchRecordZoneChangesOperation) RecordWasChangedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) SetRecordWasChangedBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) RecordZoneChangeTokensUpdatedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) FetchRecordZoneChangesCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) ConfigurationsByRecordZoneID() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) RecordChangedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) SetRecordChangedBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) RecordWithIDWasDeletedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) RecordZoneFetchCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) FetchAllChanges() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) RecordZoneIDs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) SetRecordZoneIDs() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) SetConfigurationsByRecordZoneID() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) SetRecordZoneChangeTokensUpdatedBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) SetRecordZoneFetchCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) SetFetchRecordZoneChangesCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) SetFetchAllChanges() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) SetRecordWithIDWasDeletedBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOperation) InitWithRecordZoneIDs() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


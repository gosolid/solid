//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CloudKit.CKFetchDatabaseChangesOperation : CloudKit.CKDatabaseOperation
*/

type CKFetchDatabaseChangesOperation struct {
  *CKDatabaseOperation

}

func (r *CKFetchDatabaseChangesOperation) FetchDatabaseChangesCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) SetPreviousServerChangeToken() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) FetchAllChanges() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) SetRecordZoneWithIDChangedBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) SetChangeTokenUpdatedBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) RecordZoneWithIDWasDeletedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) SetRecordZoneWithIDWasDeletedBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) PreviousServerChangeToken() (*CKServerChangeToken, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) ResultsLimit() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) SetResultsLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) RecordZoneWithIDChangedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) RecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) SetFetchDatabaseChangesCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) InitWithPreviousServerChangeToken() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) SetFetchAllChanges() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) SetRecordZoneWithIDWasPurgedBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) RecordZoneWithIDWasPurgedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) SetRecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchDatabaseChangesOperation) ChangeTokenUpdatedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}


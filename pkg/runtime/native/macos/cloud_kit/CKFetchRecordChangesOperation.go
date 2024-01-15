//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKFetchRecordChangesOperation : CloudKit.CKDatabaseOperation
*/

type CKFetchRecordChangesOperation struct {
  *CKDatabaseOperation

}

func (r *CKFetchRecordChangesOperation) RecordZoneID() (*CKRecordZoneID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) FetchRecordChangesCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) SetFetchRecordChangesCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) MoreComing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) ResultsLimit() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) RecordChangedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) RecordWithIDWasDeletedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) DesiredKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) SetDesiredKeys() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) SetRecordChangedBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) SetRecordWithIDWasDeletedBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) InitWithRecordZoneID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) SetRecordZoneID() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) PreviousServerChangeToken() (*CKServerChangeToken, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) SetResultsLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordChangesOperation) SetPreviousServerChangeToken() error {
  return fmt.Errorf("unimplemented")
}


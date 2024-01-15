//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKFetchRecordsOperation : CloudKit.CKDatabaseOperation
*/

type CKFetchRecordsOperation struct {
  *CKDatabaseOperation

}

func (r *CKFetchRecordsOperation) SetFetchRecordsCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordsOperation) InitWithRecordIDs() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordsOperation) SetPerRecordProgressBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordsOperation) PerRecordCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordsOperation) FetchRecordsCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordsOperation) FetchCurrentUserRecordOperation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordsOperation) SetRecordIDs() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordsOperation) PerRecordProgressBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordsOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordsOperation) RecordIDs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordsOperation) DesiredKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordsOperation) SetDesiredKeys() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordsOperation) SetPerRecordCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}


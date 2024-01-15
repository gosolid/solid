//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKQueryOperation : CloudKit.CKDatabaseOperation
*/

type CKQueryOperation struct {
  *CKDatabaseOperation

}

func (r *CKQueryOperation) Cursor() (*CKQueryCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) SetCursor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) ZoneID() (*CKRecordZoneID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) SetZoneID() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) SetResultsLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) SetRecordFetchedBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) InitWithQuery() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) InitWithCursor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) SetRecordMatchedBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) SetDesiredKeys() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) RecordFetchedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) QueryCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) SetQueryCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) Query() (*CKQuery, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) ResultsLimit() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) DesiredKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) RecordMatchedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQueryOperation) SetQuery() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKFetchRecordZonesOperation : CloudKit.CKDatabaseOperation
*/

type CKFetchRecordZonesOperation struct {
  *CKDatabaseOperation

}

func (r *CKFetchRecordZonesOperation) FetchAllRecordZonesOperation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZonesOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZonesOperation) InitWithRecordZoneIDs() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZonesOperation) RecordZoneIDs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZonesOperation) SetRecordZoneIDs() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZonesOperation) SetPerRecordZoneCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZonesOperation) FetchRecordZonesCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZonesOperation) SetFetchRecordZonesCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZonesOperation) PerRecordZoneCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}


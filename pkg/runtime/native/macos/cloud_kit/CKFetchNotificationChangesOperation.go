//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CloudKit.CKFetchNotificationChangesOperation : CloudKit.CKOperation
*/

type CKFetchNotificationChangesOperation struct {
  *CKOperation

}

func (r *CKFetchNotificationChangesOperation) InitWithPreviousServerChangeToken() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchNotificationChangesOperation) PreviousServerChangeToken() (*CKServerChangeToken, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchNotificationChangesOperation) MoreComing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKFetchNotificationChangesOperation) NotificationChangedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchNotificationChangesOperation) SetNotificationChangedBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchNotificationChangesOperation) SetFetchNotificationChangesCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchNotificationChangesOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchNotificationChangesOperation) SetPreviousServerChangeToken() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchNotificationChangesOperation) ResultsLimit() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKFetchNotificationChangesOperation) SetResultsLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchNotificationChangesOperation) FetchNotificationChangesCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}


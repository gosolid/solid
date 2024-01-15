//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreData.NSPersistentHistoryChangeRequest : CoreData.NSPersistentStoreRequest
*/

type NSPersistentHistoryChangeRequest struct {
  *NSPersistentStoreRequest

}

func (r *NSPersistentHistoryChangeRequest) DeleteHistoryBeforeDate() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChangeRequest) FetchRequest() (*NSFetchRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChangeRequest) SetFetchRequest() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChangeRequest) FetchHistoryAfterToken() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChangeRequest) FetchHistoryAfterTransaction() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChangeRequest) FetchHistoryWithFetchRequest() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChangeRequest) ResultType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChangeRequest) SetResultType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChangeRequest) Token() (*NSPersistentHistoryToken, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChangeRequest) FetchHistoryAfterDate() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChangeRequest) DeleteHistoryBeforeToken() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChangeRequest) DeleteHistoryBeforeTransaction() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


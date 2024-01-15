//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreData.NSBatchDeleteRequest : CoreData.NSPersistentStoreRequest
*/

type NSBatchDeleteRequest struct {
  *NSPersistentStoreRequest

}

func (r *NSBatchDeleteRequest) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchDeleteRequest) InitWithFetchRequest() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchDeleteRequest) InitWithObjectIDs() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchDeleteRequest) ResultType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBatchDeleteRequest) SetResultType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBatchDeleteRequest) FetchRequest() (*NSFetchRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}


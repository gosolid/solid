//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreData.NSAsynchronousFetchRequest : CoreData.NSPersistentStoreRequest
*/

type NSAsynchronousFetchRequest struct {
  *NSPersistentStoreRequest

}

func (r *NSAsynchronousFetchRequest) SetEstimatedResultCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAsynchronousFetchRequest) InitWithFetchRequest() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAsynchronousFetchRequest) FetchRequest() (*NSFetchRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAsynchronousFetchRequest) CompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAsynchronousFetchRequest) EstimatedResultCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}


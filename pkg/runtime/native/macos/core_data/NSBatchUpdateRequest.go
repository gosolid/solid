//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSBatchUpdateRequest : CoreData.NSPersistentStoreRequest
*/

type NSBatchUpdateRequest struct {
  *NSPersistentStoreRequest

}

func (r *NSBatchUpdateRequest) SetPropertiesToUpdate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBatchUpdateRequest) InitWithEntity() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchUpdateRequest) Predicate() (*foundation.NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchUpdateRequest) SetIncludesSubentities() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBatchUpdateRequest) BatchUpdateRequestWithEntityName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchUpdateRequest) EntityName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchUpdateRequest) ResultType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBatchUpdateRequest) IncludesSubentities() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBatchUpdateRequest) SetResultType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBatchUpdateRequest) PropertiesToUpdate() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchUpdateRequest) InitWithEntityName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchUpdateRequest) Entity() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchUpdateRequest) SetPredicate() error {
  return fmt.Errorf("unimplemented")
}


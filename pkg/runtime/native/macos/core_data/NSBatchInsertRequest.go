//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSBatchInsertRequest : CoreData.NSPersistentStoreRequest
*/

type NSBatchInsertRequest struct {
  *NSPersistentStoreRequest

}

func (r *NSBatchInsertRequest) InitWithEntityName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchInsertRequest) EntityName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchInsertRequest) ObjectsToInsert() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchInsertRequest) ManagedObjectHandler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchInsertRequest) SetResultType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBatchInsertRequest) DictionaryHandler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchInsertRequest) SetManagedObjectHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBatchInsertRequest) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchInsertRequest) InitWithEntity() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchInsertRequest) SetObjectsToInsert() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBatchInsertRequest) SetDictionaryHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBatchInsertRequest) ResultType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBatchInsertRequest) BatchInsertRequestWithEntityName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchInsertRequest) Entity() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}


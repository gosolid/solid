//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreData.NSCustomMigrationStage : CoreData.NSMigrationStage
*/

type NSCustomMigrationStage struct {
  *NSMigrationStage

}

func (r *NSCustomMigrationStage) WillMigrateHandler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCustomMigrationStage) SetWillMigrateHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCustomMigrationStage) DidMigrateHandler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCustomMigrationStage) SetDidMigrateHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCustomMigrationStage) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCustomMigrationStage) InitWithCurrentModelReference() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCustomMigrationStage) CurrentModel() (*NSManagedObjectModelReference, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCustomMigrationStage) NextModel() (*NSManagedObjectModelReference, error) {
  return nil, fmt.Errorf("unimplemented")
}


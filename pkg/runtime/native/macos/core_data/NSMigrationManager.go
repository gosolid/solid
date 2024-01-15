//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSMigrationManager : objc.NSObject
*/

type NSMigrationManager struct {
  *objc.NSObject

}

func (r *NSMigrationManager) InitWithSourceModel() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) DestinationInstancesForEntityMappingNamed() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) UsesStoreSpecificMigrationManager() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) MappingModel() (*NSMappingModel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) SourceModel() (*NSManagedObjectModel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) DestinationEntityForEntityMapping() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) CurrentEntityMapping() (*NSEntityMapping, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) SetUserInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) SourceEntityForEntityMapping() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) SourceInstancesForEntityMappingNamed() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) CancelMigrationWithError() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) SetUsesStoreSpecificMigrationManager() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) UserInfo() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) MigrationProgress() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) MigrateStoreFromURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) Reset() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) AssociateSourceInstance() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) DestinationModel() (*NSManagedObjectModel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) SourceContext() (*NSManagedObjectContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMigrationManager) DestinationContext() (*NSManagedObjectContext, error) {
  return nil, fmt.Errorf("unimplemented")
}


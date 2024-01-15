//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface CoreData.NSPersistentStoreCoordinator : objc.NSObject
*/

type NSPersistentStoreCoordinator struct {
  *objc.NSObject
  *foundation.NSLocking
}

func (r *NSPersistentStoreCoordinator) ImportStoreWithIdentifier() (*NSPersistentStore, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) PerformBlockAndWait() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) FinishDeferredLightweightMigration() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) TryLock() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) ManagedObjectModel() (*NSManagedObjectModel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) PersistentStores() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) RegisteredStoreTypes() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) InitWithManagedObjectModel() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) AddPersistentStoreWithDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) RemovePersistentStore() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) ReplacePersistentStoreAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) RemoveUbiquitousContentAndPersistentStoreAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) SetMetadata() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) ManagedObjectIDForURIRepresentation() (*NSManagedObjectID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) FinishDeferredLightweightMigrationTask() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) PersistentStoreForURL() (*NSPersistentStore, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) AddPersistentStoreWithType() (*NSPersistentStore, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) MetadataForPersistentStore() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) ExecuteRequest() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) SetURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) ElementsDerivedFromExternalRecordURL() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) MigratePersistentStore() (*NSPersistentStore, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) PerformBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) CurrentPersistentHistoryTokenFromStores() (*NSPersistentHistoryToken, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) MetadataForPersistentStoreOfType() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) DestroyPersistentStoreAtURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) Lock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) Unlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) Name() (*objc.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) URLForPersistentStore() (*core_foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) RegisterStoreClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) MetadataForPersistentStoreWithURL() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreCoordinator) SetName() error {
  return fmt.Errorf("unimplemented")
}


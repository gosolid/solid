//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface CoreData.NSPersistentContainer : objc.NSObject
*/

type NSPersistentContainer struct {
  *objc.NSObject

}

func (r *NSPersistentContainer) ViewContext() (*NSManagedObjectContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentContainer) ManagedObjectModel() (*NSManagedObjectModel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentContainer) PersistentStoreDescriptions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentContainer) SetPersistentStoreDescriptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentContainer) InitWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentContainer) LoadPersistentStoresWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentContainer) NewBackgroundContext() (*NSManagedObjectContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentContainer) Name() (*objc.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentContainer) PersistentContainerWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentContainer) DefaultDirectoryURL() (*core_foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentContainer) PerformBackgroundTask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentContainer) PersistentStoreCoordinator() (*NSPersistentStoreCoordinator, error) {
  return nil, fmt.Errorf("unimplemented")
}


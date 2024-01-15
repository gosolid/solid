//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_data"
)

/*
interface AppKit.NSPersistentDocument : AppKit.NSDocument
*/

type NSPersistentDocument struct {
  *NSDocument

}

func (r *NSPersistentDocument) ReadFromURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentDocument) RevertToContentsOfURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentDocument) ManagedObjectContext() (*core_data.NSManagedObjectContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentDocument) SetManagedObjectContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentDocument) ManagedObjectModel() (*NSManagedObjectModel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentDocument) ConfigurePersistentStoreCoordinatorForURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentDocument) PersistentStoreTypeForFileType() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentDocument) WriteToURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


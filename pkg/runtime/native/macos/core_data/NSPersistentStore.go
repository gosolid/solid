//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CoreData.NSPersistentStore : objc.NSObject
*/

type NSPersistentStore struct {
  *objc.NSObject

}

func (r *NSPersistentStore) URL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) Type() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) IsReadOnly() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) SetReadOnly() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) LoadMetadata() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) Options() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) SetURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) PersistentStoreCoordinator() (*NSPersistentStoreCoordinator, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) ConfigurationName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) WillRemoveFromPersistentStoreCoordinator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) Identifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) Metadata() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) MetadataForPersistentStoreWithURL() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) DidAddToPersistentStoreCoordinator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) InitWithPersistentStoreCoordinator() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) SetIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) CoreSpotlightExporter() (*NSCoreDataCoreSpotlightDelegate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) SetMetadata() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStore) MigrationManagerClass() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSIncrementalStore : CoreData.NSPersistentStore
*/

type NSIncrementalStore struct {
  *NSPersistentStore

}

func (r *NSIncrementalStore) NewValuesForObjectWithID() (*NSIncrementalStoreNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIncrementalStore) NewValueForRelationship() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIncrementalStore) ReferenceObjectForObjectID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIncrementalStore) NewObjectIDForEntity() (*NSManagedObjectID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIncrementalStore) LoadMetadata() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSIncrementalStore) ExecuteRequest() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIncrementalStore) IdentifierForNewStoreAtURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIncrementalStore) ObtainPermanentIDsForObjects() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIncrementalStore) ManagedObjectContextDidRegisterObjectsWithIDs() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSIncrementalStore) ManagedObjectContextDidUnregisterObjectsWithIDs() error {
  return fmt.Errorf("unimplemented")
}


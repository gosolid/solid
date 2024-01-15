//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreData.NSAtomicStore : CoreData.NSPersistentStore
*/

type NSAtomicStore struct {
  *NSPersistentStore

}

func (r *NSAtomicStore) NewCacheNodeForManagedObject() (*NSAtomicStoreCacheNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAtomicStore) WillRemoveCacheNodes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAtomicStore) CacheNodeForObjectID() (*NSAtomicStoreCacheNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAtomicStore) ObjectIDForEntity() (*NSManagedObjectID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAtomicStore) NewReferenceObjectForManagedObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAtomicStore) Save() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAtomicStore) Load() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAtomicStore) UpdateCacheNode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAtomicStore) CacheNodes() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAtomicStore) AddCacheNodes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAtomicStore) ReferenceObjectForObjectID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAtomicStore) InitWithPersistentStoreCoordinator() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


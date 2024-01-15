//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreData.NSPersistentHistoryChange : objc.NSObject
*/

type NSPersistentHistoryChange struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSPersistentHistoryChange) Transaction() (*NSPersistentHistoryTransaction, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChange) UpdatedProperties() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChange) EntityDescription() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChange) ChangeID() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChange) ChangeType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChange) Tombstone() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChange) EntityDescriptionWithContext() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChange) FetchRequest() (*NSFetchRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryChange) ChangedObjectID() (*NSManagedObjectID, error) {
  return nil, fmt.Errorf("unimplemented")
}


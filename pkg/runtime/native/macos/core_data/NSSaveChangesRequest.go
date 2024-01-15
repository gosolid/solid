//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSSaveChangesRequest : CoreData.NSPersistentStoreRequest
*/

type NSSaveChangesRequest struct {
  *NSPersistentStoreRequest

}

func (r *NSSaveChangesRequest) InitWithInsertedObjects() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSaveChangesRequest) InsertedObjects() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSaveChangesRequest) UpdatedObjects() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSaveChangesRequest) DeletedObjects() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSaveChangesRequest) LockedObjects() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}


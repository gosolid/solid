//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreData.NSManagedObjectID : objc.NSObject
*/

type NSManagedObjectID struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSManagedObjectID) URIRepresentation() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectID) Entity() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectID) PersistentStore() (*NSPersistentStore, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectID) IsTemporaryID() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


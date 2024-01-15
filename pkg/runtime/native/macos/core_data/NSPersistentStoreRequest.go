//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreData.NSPersistentStoreRequest : objc.NSObject
*/

type NSPersistentStoreRequest struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSPersistentStoreRequest) AffectedStores() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreRequest) SetAffectedStores() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreRequest) RequestType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


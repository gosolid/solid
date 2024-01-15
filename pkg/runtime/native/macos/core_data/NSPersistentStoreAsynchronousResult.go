//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSPersistentStoreAsynchronousResult : CoreData.NSPersistentStoreResult
*/

type NSPersistentStoreAsynchronousResult struct {
  *NSPersistentStoreResult

}

func (r *NSPersistentStoreAsynchronousResult) Cancel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreAsynchronousResult) ManagedObjectContext() (*NSManagedObjectContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreAsynchronousResult) OperationError() (*foundation.NSError, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreAsynchronousResult) Progress() (*foundation.NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}


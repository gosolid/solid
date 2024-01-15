//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreData.NSPersistentHistoryResult : CoreData.NSPersistentStoreResult
*/

type NSPersistentHistoryResult struct {
  *NSPersistentStoreResult

}

func (r *NSPersistentHistoryResult) Result() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryResult) ResultType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


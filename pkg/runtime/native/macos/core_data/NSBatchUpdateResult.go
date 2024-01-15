//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreData.NSBatchUpdateResult : CoreData.NSPersistentStoreResult
*/

type NSBatchUpdateResult struct {
  *NSPersistentStoreResult

}

func (r *NSBatchUpdateResult) Result() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchUpdateResult) ResultType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


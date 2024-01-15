//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreData.NSBatchInsertResult : CoreData.NSPersistentStoreResult
*/

type NSBatchInsertResult struct {
  *NSPersistentStoreResult

}

func (r *NSBatchInsertResult) Result() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchInsertResult) ResultType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


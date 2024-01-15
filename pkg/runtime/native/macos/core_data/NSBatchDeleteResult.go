//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreData.NSBatchDeleteResult : CoreData.NSPersistentStoreResult
*/

type NSBatchDeleteResult struct {
  *NSPersistentStoreResult

}

func (r *NSBatchDeleteResult) Result() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBatchDeleteResult) ResultType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


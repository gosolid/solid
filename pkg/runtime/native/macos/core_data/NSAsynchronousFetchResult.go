//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSAsynchronousFetchResult : CoreData.NSPersistentStoreAsynchronousResult
*/

type NSAsynchronousFetchResult struct {
  *NSPersistentStoreAsynchronousResult

}

func (r *NSAsynchronousFetchResult) FetchRequest() (*NSAsynchronousFetchRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAsynchronousFetchResult) FinalResult() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


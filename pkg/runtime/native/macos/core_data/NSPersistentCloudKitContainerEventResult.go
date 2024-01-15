//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreData.NSPersistentCloudKitContainerEventResult : CoreData.NSPersistentStoreResult
*/

type NSPersistentCloudKitContainerEventResult struct {
  *NSPersistentStoreResult

}

func (r *NSPersistentCloudKitContainerEventResult) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEventResult) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEventResult) Result() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEventResult) ResultType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


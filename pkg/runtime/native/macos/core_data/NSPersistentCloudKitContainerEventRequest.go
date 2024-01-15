//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreData.NSPersistentCloudKitContainerEventRequest : CoreData.NSPersistentStoreRequest
*/

type NSPersistentCloudKitContainerEventRequest struct {
  *NSPersistentStoreRequest

}

func (r *NSPersistentCloudKitContainerEventRequest) SetResultType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEventRequest) FetchEventsAfterDate() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEventRequest) FetchEventsAfterEvent() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEventRequest) FetchEventsMatchingFetchRequest() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEventRequest) FetchRequestForEvents() (*NSFetchRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEventRequest) ResultType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


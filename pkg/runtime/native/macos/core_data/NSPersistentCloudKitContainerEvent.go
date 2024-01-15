//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreData.NSPersistentCloudKitContainerEvent : objc.NSObject
*/

type NSPersistentCloudKitContainerEvent struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSPersistentCloudKitContainerEvent) Error() (*foundation.NSError, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEvent) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEvent) Identifier() (*foundation.NSUUID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEvent) StoreIdentifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEvent) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEvent) EndDate() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEvent) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEvent) StartDate() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerEvent) Succeeded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CoreData.NSPersistentCloudKitContainerOptions : objc.NSObject
*/

type NSPersistentCloudKitContainerOptions struct {
  *objc.NSObject

}

func (r *NSPersistentCloudKitContainerOptions) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerOptions) InitWithContainerIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerOptions) ContainerIdentifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerOptions) DatabaseScope() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainerOptions) SetDatabaseScope() error {
  return fmt.Errorf("unimplemented")
}


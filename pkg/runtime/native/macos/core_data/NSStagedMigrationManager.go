//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CoreData.NSStagedMigrationManager : objc.NSObject
*/

type NSStagedMigrationManager struct {
  *objc.NSObject

}

func (r *NSStagedMigrationManager) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStagedMigrationManager) InitWithMigrationStages() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStagedMigrationManager) Stages() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStagedMigrationManager) Container() (*NSPersistentContainer, error) {
  return nil, fmt.Errorf("unimplemented")
}


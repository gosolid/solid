//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSMigrationStage : objc.NSObject
*/

type NSMigrationStage struct {
  *objc.NSObject

}

func (r *NSMigrationStage) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMigrationStage) SetLabel() error {
  return fmt.Errorf("unimplemented")
}


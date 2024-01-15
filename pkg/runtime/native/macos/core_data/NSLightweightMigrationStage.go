//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSLightweightMigrationStage : CoreData.NSMigrationStage
*/

type NSLightweightMigrationStage struct {
  *NSMigrationStage

}

func (r *NSLightweightMigrationStage) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLightweightMigrationStage) InitWithVersionChecksums() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLightweightMigrationStage) VersionChecksums() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CoreData.NSEntityMigrationPolicy : objc.NSObject
*/

type NSEntityMigrationPolicy struct {
  *objc.NSObject

}

func (r *NSEntityMigrationPolicy) CreateDestinationInstancesForSourceInstance() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEntityMigrationPolicy) EndInstanceCreationForEntityMapping() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEntityMigrationPolicy) CreateRelationshipsForDestinationInstance() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEntityMigrationPolicy) EndRelationshipCreationForEntityMapping() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEntityMigrationPolicy) PerformCustomValidationForEntityMapping() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEntityMigrationPolicy) EndEntityMapping() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEntityMigrationPolicy) BeginEntityMapping() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


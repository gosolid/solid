//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSConstraintConflict : objc.NSObject
*/

type NSConstraintConflict struct {
  *objc.NSObject

}

func (r *NSConstraintConflict) ConflictingSnapshots() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConstraintConflict) InitWithConstraint() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConstraintConflict) Constraint() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConstraintConflict) ConstraintValues() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConstraintConflict) DatabaseObject() (*NSManagedObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConstraintConflict) DatabaseSnapshot() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConstraintConflict) ConflictingObjects() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


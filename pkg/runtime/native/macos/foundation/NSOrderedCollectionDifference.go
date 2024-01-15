//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSOrderedCollectionDifference : objc.NSObject
*/

type NSOrderedCollectionDifference struct {
  *objc.NSObject
  *NSFastEnumeration
}

func (r *NSOrderedCollectionDifference) HasChanges() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionDifference) InitWithChanges() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionDifference) InitWithInsertIndexes() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionDifference) DifferenceByTransformingChangesWithBlock() (*NSOrderedCollectionDifference, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionDifference) InverseDifference() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionDifference) Insertions() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionDifference) Removals() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


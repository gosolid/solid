//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSOrderedCollectionDifference : objc.NSObject
*/

type NSOrderedCollectionDifference struct {
  *objc.NSObject

}

func (r *NSOrderedCollectionDifference) Removals() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionDifference) HasChanges() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionDifference) InitWithChanges() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionDifference) InitWithInsertIndexes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionDifference) DifferenceByTransformingChangesWithBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionDifference) InverseDifference() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionDifference) Insertions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


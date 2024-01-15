//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CoreData.NSIncrementalStoreNode : objc.NSObject
*/

type NSIncrementalStoreNode struct {
  *objc.NSObject

}

func (r *NSIncrementalStoreNode) InitWithObjectID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIncrementalStoreNode) UpdateWithValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSIncrementalStoreNode) ValueForPropertyDescription() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIncrementalStoreNode) ObjectID() (*NSManagedObjectID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIncrementalStoreNode) Version() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}


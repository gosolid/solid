//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSAtomicStoreCacheNode : objc.NSObject
*/

type NSAtomicStoreCacheNode struct {
  *objc.NSObject

}

func (r *NSAtomicStoreCacheNode) ObjectID() (*NSManagedObjectID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAtomicStoreCacheNode) PropertyCache() (*foundation.NSMutableDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAtomicStoreCacheNode) SetPropertyCache() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAtomicStoreCacheNode) InitWithObjectID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAtomicStoreCacheNode) ValueForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAtomicStoreCacheNode) SetValue() error {
  return fmt.Errorf("unimplemented")
}


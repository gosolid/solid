//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSOrderedCollectionChange : objc.NSObject
*/

type NSOrderedCollectionChange struct {
  *objc.NSObject

}

func (r *NSOrderedCollectionChange) ChangeType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionChange) Index() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionChange) AssociatedIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionChange) ChangeWithObject() (*NSOrderedCollectionChange, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionChange) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionChange) InitWithObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionChange) Object() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


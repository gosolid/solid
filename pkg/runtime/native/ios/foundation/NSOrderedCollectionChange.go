//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSOrderedCollectionChange : objc.NSObject
*/

type NSOrderedCollectionChange struct {
  *objc.NSObject

}

func (r *NSOrderedCollectionChange) ChangeWithObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionChange) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionChange) InitWithObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionChange) Object() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionChange) ChangeType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionChange) Index() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOrderedCollectionChange) AssociatedIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


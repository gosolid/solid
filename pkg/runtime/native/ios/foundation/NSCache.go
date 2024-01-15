//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSCache : objc.NSObject
*/

type NSCache struct {
  *objc.NSObject

}

func (r *NSCache) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) CountLimit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) EvictsObjectsWithDiscardedContent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) ObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) RemoveAllObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) SetCountLimit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) RemoveObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) TotalCostLimit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) SetTotalCostLimit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) SetEvictsObjectsWithDiscardedContent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


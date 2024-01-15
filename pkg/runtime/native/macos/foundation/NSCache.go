//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSCache : objc.NSObject
*/

type NSCache struct {
  *objc.NSObject

}

func (r *NSCache) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCache) SetEvictsObjectsWithDiscardedContent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCache) ObjectForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCache) TotalCostLimit() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCache) SetTotalCostLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCache) CountLimit() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCache) EvictsObjectsWithDiscardedContent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCache) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCache) SetCountLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCache) SetObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCache) RemoveObjectForKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCache) RemoveAllObjects() error {
  return fmt.Errorf("unimplemented")
}


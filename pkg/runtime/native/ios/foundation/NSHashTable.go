//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSHashTable : objc.NSObject
*/

type NSHashTable struct {
  *objc.NSObject

}

func (r *NSHashTable) PointerFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) Count() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) AnyObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) ObjectEnumerator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) MinusHashTable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) RemoveObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) IsEqualToHashTable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) InitWithOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) Member() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) AddObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) RemoveAllObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) ContainsObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) IntersectHashTable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) UnionHashTable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) SetRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) InitWithPointerFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) WeakObjectsHashTable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) IntersectsHashTable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) IsSubsetOfHashTable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) AllObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) HashTableWithOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) HashTableWithWeakObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


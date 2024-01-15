//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSHashTable : objc.NSObject
*/

type NSHashTable struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
  *NSFastEnumeration
}

func (r *NSHashTable) PointerFunctions() (*NSPointerFunctions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) HashTableWithWeakObjects() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) ObjectEnumerator() (*NSEnumerator, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) RemoveObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHashTable) IntersectsHashTable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) IsSubsetOfHashTable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) UnionHashTable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHashTable) MinusHashTable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHashTable) Count() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) InitWithOptions() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) HashTableWithOptions() (*NSHashTable, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) WeakObjectsHashTable() (*NSHashTable, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) Member() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) RemoveAllObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHashTable) IsEqualToHashTable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) SetRepresentation() (*NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) ContainsObject() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) IntersectHashTable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHashTable) AllObjects() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) AnyObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) InitWithPointerFunctions() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHashTable) AddObject() error {
  return fmt.Errorf("unimplemented")
}


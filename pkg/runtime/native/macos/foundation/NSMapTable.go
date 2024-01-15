//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSMapTable : objc.NSObject
*/

type NSMapTable struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
  *NSFastEnumeration
}

func (r *NSMapTable) DictionaryRepresentation() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) KeyPointerFunctions() (*NSPointerFunctions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) InitWithKeyOptions() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) MapTableWithKeyOptions() (*NSMapTable, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) MapTableWithWeakToWeakObjects() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) StrongToStrongObjectsMapTable() (*NSMapTable, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) WeakToStrongObjectsMapTable() (*NSMapTable, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) WeakToWeakObjectsMapTable() (*NSMapTable, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) Count() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) MapTableWithStrongToStrongObjects() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) StrongToWeakObjectsMapTable() (*NSMapTable, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) RemoveObjectForKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMapTable) KeyEnumerator() (*NSEnumerator, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) MapTableWithWeakToStrongObjects() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) ObjectForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) RemoveAllObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMapTable) ValuePointerFunctions() (*NSPointerFunctions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) InitWithKeyPointerFunctions() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) MapTableWithStrongToWeakObjects() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) SetObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMapTable) ObjectEnumerator() (*NSEnumerator, error) {
  return nil, fmt.Errorf("unimplemented")
}


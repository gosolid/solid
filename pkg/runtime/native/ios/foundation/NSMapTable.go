//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSMapTable : objc.NSObject
*/

type NSMapTable struct {
  *objc.NSObject

}

func (r *NSMapTable) MapTableWithKeyOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) MapTableWithStrongToWeakObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) MapTableWithWeakToWeakObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) WeakToStrongObjectsMapTable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) ObjectEnumerator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) KeyPointerFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) InitWithKeyOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) MapTableWithWeakToStrongObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) StrongToStrongObjectsMapTable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) WeakToWeakObjectsMapTable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) RemoveAllObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) ValuePointerFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) StrongToWeakObjectsMapTable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) RemoveObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) InitWithKeyPointerFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) MapTableWithStrongToStrongObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) ObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) KeyEnumerator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) DictionaryRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMapTable) Count() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


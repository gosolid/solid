//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSPointerArray : objc.NSObject
*/

type NSPointerArray struct {
  *objc.NSObject

}

func (r *NSPointerArray) PointerAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) SetCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) InitWithOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) ReplacePointerAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) PointerFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) PointerArrayWithPointerFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) AddPointer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) PointerArrayWithOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) RemovePointerAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) InsertPointer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) Compact() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) Count() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) InitWithPointerFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


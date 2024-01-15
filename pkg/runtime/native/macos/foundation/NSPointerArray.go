//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSPointerArray : objc.NSObject
*/

type NSPointerArray struct {
  *objc.NSObject
  *NSFastEnumeration
  *NSCopying
  *NSSecureCoding
}

func (r *NSPointerArray) PointerArrayWithOptions() (*NSPointerArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) Compact() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) InitWithOptions() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) PointerAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) RemovePointerAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) ReplacePointerAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) PointerFunctions() (*NSPointerFunctions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) SetCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) InitWithPointerFunctions() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) PointerArrayWithPointerFunctions() (*NSPointerArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) Count() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) AddPointer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPointerArray) InsertPointer() error {
  return fmt.Errorf("unimplemented")
}


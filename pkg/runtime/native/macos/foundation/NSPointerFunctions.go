//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSPointerFunctions : objc.NSObject
*/

type NSPointerFunctions struct {
  *objc.NSObject
  *NSCopying
}

func (r *NSPointerFunctions) InitWithOptions() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetHashFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SizeFunction() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) AcquireFunction() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) UsesWeakReadAndWriteBarriers() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetUsesWeakReadAndWriteBarriers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) PointerFunctionsWithOptions() (*NSPointerFunctions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) HashFunction() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetIsEqualFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) RelinquishFunction() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetRelinquishFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetAcquireFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetUsesStrongWriteBarrier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) DescriptionFunction() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) UsesStrongWriteBarrier() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) IsEqualFunction() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetSizeFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetDescriptionFunction() error {
  return fmt.Errorf("unimplemented")
}


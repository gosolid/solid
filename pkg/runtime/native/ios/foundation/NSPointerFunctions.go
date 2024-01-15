//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSPointerFunctions : objc.NSObject
*/

type NSPointerFunctions struct {
  *objc.NSObject

}

func (r *NSPointerFunctions) SetUsesStrongWriteBarrier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetDescriptionFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetAcquireFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetIsEqualFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SizeFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) RelinquishFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetRelinquishFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) InitWithOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) HashFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) DescriptionFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) AcquireFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) UsesStrongWriteBarrier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) UsesWeakReadAndWriteBarriers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) PointerFunctionsWithOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetSizeFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetUsesWeakReadAndWriteBarriers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) SetHashFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPointerFunctions) IsEqualFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


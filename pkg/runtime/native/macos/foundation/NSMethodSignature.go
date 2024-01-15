//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSMethodSignature : objc.NSObject
*/

type NSMethodSignature struct {
  *objc.NSObject

}

func (r *NSMethodSignature) MethodReturnType() (byte, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMethodSignature) MethodReturnLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMethodSignature) SignatureWithObjCTypes() (*NSMethodSignature, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMethodSignature) GetArgumentTypeAtIndex() (byte, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMethodSignature) IsOneway() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMethodSignature) NumberOfArguments() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMethodSignature) FrameLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}


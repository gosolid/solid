//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSMethodSignature : objc.NSObject
*/

type NSMethodSignature struct {
  *objc.NSObject

}

func (r *NSMethodSignature) IsOneway() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMethodSignature) NumberOfArguments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMethodSignature) FrameLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMethodSignature) MethodReturnType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMethodSignature) MethodReturnLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMethodSignature) SignatureWithObjCTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMethodSignature) GetArgumentTypeAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


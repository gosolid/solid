//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSInvocation : objc.NSObject
*/

type NSInvocation struct {
  *objc.NSObject

}

func (r *NSInvocation) GetReturnValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) GetArgument() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) InvokeWithTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) ArgumentsRetained() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) Target() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) Selector() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) InvocationWithMethodSignature() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) RetainArguments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) Invoke() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) MethodSignature() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) SetSelector() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) SetTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) SetReturnValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) SetArgument() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) InvokeUsingIMP() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


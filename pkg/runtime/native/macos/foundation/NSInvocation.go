//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSInvocation : objc.NSObject
*/

type NSInvocation struct {
  *objc.NSObject

}

func (r *NSInvocation) SetSelector() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInvocation) Target() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) RetainArguments() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInvocation) SetArgument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInvocation) Invoke() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInvocation) InvokeUsingIMP() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInvocation) SetTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInvocation) InvocationWithMethodSignature() (*NSInvocation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) GetArgument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInvocation) MethodSignature() (*NSMethodSignature, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) Selector() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) SetReturnValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInvocation) InvokeWithTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInvocation) ArgumentsRetained() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSInvocation) GetReturnValue() error {
  return fmt.Errorf("unimplemented")
}


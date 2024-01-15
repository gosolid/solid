//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLLibrary
*/

type MTLLibrary struct {

}

func (r *MTLLibrary) NewFunctionWithDescriptor() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLLibrary) NewIntersectionFunctionWithDescriptor() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLLibrary) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLibrary) FunctionNames() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLibrary) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLLibrary) InstallName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLibrary) NewFunctionWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLibrary) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLLibrary) SetLabel() error {
  return fmt.Errorf("unimplemented")
}


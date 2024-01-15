//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLStencilDescriptor : objc.NSObject
*/

type MTLStencilDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLStencilDescriptor) SetStencilCompareFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) StencilFailureOperation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) DepthStencilPassOperation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) SetWriteMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) SetReadMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) WriteMask() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) StencilCompareFunction() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) SetStencilFailureOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) DepthFailureOperation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) SetDepthFailureOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) SetDepthStencilPassOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) ReadMask() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}


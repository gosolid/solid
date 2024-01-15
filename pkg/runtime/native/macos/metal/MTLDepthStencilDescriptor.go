//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLDepthStencilDescriptor : objc.NSObject
*/

type MTLDepthStencilDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLDepthStencilDescriptor) DepthCompareFunction() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) SetDepthCompareFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) SetDepthWriteEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) SetBackFaceStencil() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) IsDepthWriteEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) FrontFaceStencil() (*MTLStencilDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) SetFrontFaceStencil() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) BackFaceStencil() (*MTLStencilDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}


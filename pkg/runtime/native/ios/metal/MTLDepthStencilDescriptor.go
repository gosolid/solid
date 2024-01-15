//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLDepthStencilDescriptor : objc.NSObject
*/

type MTLDepthStencilDescriptor struct {
  *objc.NSObject

}

func (r *MTLDepthStencilDescriptor) SetDepthCompareFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) SetFrontFaceStencil() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) BackFaceStencil() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) SetLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) DepthCompareFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) SetDepthWriteEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) FrontFaceStencil() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) SetBackFaceStencil() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) Label() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilDescriptor) IsDepthWriteEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


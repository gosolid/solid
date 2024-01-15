//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLStencilDescriptor : objc.NSObject
*/

type MTLStencilDescriptor struct {
  *objc.NSObject

}

func (r *MTLStencilDescriptor) SetStencilCompareFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) StencilFailureOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) SetStencilFailureOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) DepthFailureOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) ReadMask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) SetReadMask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) SetWriteMask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) StencilCompareFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) SetDepthFailureOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) DepthStencilPassOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) SetDepthStencilPassOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStencilDescriptor) WriteMask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLRenderPassAttachmentDescriptor : objc.NSObject
*/

type MTLRenderPassAttachmentDescriptor struct {
  *objc.NSObject

}

func (r *MTLRenderPassAttachmentDescriptor) ResolveLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetResolveSlice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetLoadAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetDepthPlane() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) ResolveTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetResolveDepthPlane() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) StoreAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) Slice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetResolveTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetSlice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) ResolveDepthPlane() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) LoadAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetStoreAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) Texture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) Level() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetResolveLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) ResolveSlice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) StoreActionOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetStoreActionOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) DepthPlane() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


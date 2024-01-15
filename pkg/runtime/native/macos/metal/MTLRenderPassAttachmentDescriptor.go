//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLRenderPassAttachmentDescriptor : objc.NSObject
*/

type MTLRenderPassAttachmentDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLRenderPassAttachmentDescriptor) SetLoadAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetStoreAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetSlice() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetResolveSlice() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) DepthPlane() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) ResolveLevel() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetResolveLevel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) LoadAction() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) Texture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetLevel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) StoreActionOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) Slice() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetDepthPlane() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) ResolveTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetResolveTexture() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) ResolveSlice() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) ResolveDepthPlane() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetResolveDepthPlane() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) StoreAction() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetTexture() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) Level() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassAttachmentDescriptor) SetStoreActionOptions() error {
  return fmt.Errorf("unimplemented")
}


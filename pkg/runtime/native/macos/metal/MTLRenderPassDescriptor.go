//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLRenderPassDescriptor : objc.NSObject
*/

type MTLRenderPassDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLRenderPassDescriptor) SetVisibilityResultBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) RasterizationRateMap() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) VisibilityResultBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) ThreadgroupMemoryLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) TileWidth() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetTileHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) RenderTargetWidth() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetRenderTargetHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetRasterizationRateMap() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetStencilAttachment() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetThreadgroupMemoryLength() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetRenderTargetArrayLength() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetImageblockSampleLength() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetSamplePositions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetDepthAttachment() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetTileWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) RenderTargetHeight() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) GetSamplePositions() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) DepthAttachment() (*MTLRenderPassDepthAttachmentDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) StencilAttachment() (*MTLRenderPassStencilAttachmentDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) RenderTargetArrayLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) RenderPassDescriptor() (*MTLRenderPassDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) DefaultRasterSampleCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetDefaultRasterSampleCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetRenderTargetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) ColorAttachments() (*MTLRenderPassColorAttachmentDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) TileHeight() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SampleBufferAttachments() (*MTLRenderPassSampleBufferAttachmentDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) ImageblockSampleLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}


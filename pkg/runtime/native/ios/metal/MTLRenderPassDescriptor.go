//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLRenderPassDescriptor : objc.NSObject
*/

type MTLRenderPassDescriptor struct {
  *objc.NSObject

}

func (r *MTLRenderPassDescriptor) SetRenderTargetArrayLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetImageblockSampleLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetRenderTargetWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetTileWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) RenderTargetWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetRenderTargetHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetSamplePositions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) DepthAttachment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetDepthAttachment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetStencilAttachment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) DefaultRasterSampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) ImageblockSampleLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) TileWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetTileHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) RasterizationRateMap() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) ColorAttachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetVisibilityResultBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) StencilAttachment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) RenderTargetArrayLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) RenderTargetHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetRasterizationRateMap() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) RenderPassDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) VisibilityResultBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) ThreadgroupMemoryLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetThreadgroupMemoryLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) TileHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) GetSamplePositions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SetDefaultRasterSampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDescriptor) SampleBufferAttachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


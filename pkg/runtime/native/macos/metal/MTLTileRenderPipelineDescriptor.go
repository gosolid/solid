//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLTileRenderPipelineDescriptor : objc.NSObject
*/

type MTLTileRenderPipelineDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLTileRenderPipelineDescriptor) SetLinkedFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SupportAddingBinaryFunctions() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) MaxCallStackDepth() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetMaxCallStackDepth() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) Reset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) RasterSampleCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetMaxTotalThreadsPerThreadgroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetPreloadedLibraries() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetTileFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) LinkedFunctions() (*MTLLinkedFunctions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) PreloadedLibraries() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) TileFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetThreadgroupSizeMatchesTileSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetBinaryArchives() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) TileBuffers() (*MTLPipelineBufferDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) MaxTotalThreadsPerThreadgroup() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) BinaryArchives() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetSupportAddingBinaryFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetRasterSampleCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) ColorAttachments() (*MTLTileRenderPipelineColorAttachmentDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) ThreadgroupSizeMatchesTileSize() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


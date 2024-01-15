//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLTileRenderPipelineDescriptor : objc.NSObject
*/

type MTLTileRenderPipelineDescriptor struct {
  *objc.NSObject

}

func (r *MTLTileRenderPipelineDescriptor) SetLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetRasterSampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SupportAddingBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) TileFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) RasterSampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) ThreadgroupSizeMatchesTileSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetMaxTotalThreadsPerThreadgroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetBinaryArchives() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) PreloadedLibraries() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetMaxCallStackDepth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) Reset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) TileBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) LinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetSupportAddingBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetLinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) Label() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetTileFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) ColorAttachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetThreadgroupSizeMatchesTileSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) MaxTotalThreadsPerThreadgroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) BinaryArchives() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) SetPreloadedLibraries() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineDescriptor) MaxCallStackDepth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


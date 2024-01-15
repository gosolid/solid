//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLRenderPipelineState
*/

type MTLRenderPipelineState struct {

}

func (r *MTLRenderPipelineState) MaxTotalThreadgroupsPerMeshGrid() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) GpuResourceID() (MTLResourceID, error) {
  return MTLResourceID{}, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) ImageblockMemoryLengthForDimensions() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) MaxTotalThreadsPerMeshThreadgroup() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) MeshThreadExecutionWidth() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) ThreadgroupSizeMatchesTileSize() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) ImageblockSampleLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) MaxTotalThreadsPerObjectThreadgroup() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) ObjectThreadExecutionWidth() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) NewRenderPipelineStateWithAdditionalBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) MaxTotalThreadsPerThreadgroup() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) SupportIndirectCommandBuffers() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) FunctionHandleWithFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) NewVisibleFunctionTableWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineState) NewIntersectionFunctionTableWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


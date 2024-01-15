//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLComputePipelineState
*/

type MTLComputePipelineState struct {

}

func (r *MTLComputePipelineState) NewIntersectionFunctionTableWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineState) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineState) MaxTotalThreadsPerThreadgroup() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineState) SupportIndirectCommandBuffers() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineState) GpuResourceID() (MTLResourceID, error) {
  return MTLResourceID{}, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineState) FunctionHandleWithFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineState) NewComputePipelineStateWithAdditionalBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineState) NewVisibleFunctionTableWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineState) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineState) ThreadExecutionWidth() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineState) StaticThreadgroupMemoryLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineState) ImageblockMemoryLengthForDimensions() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}


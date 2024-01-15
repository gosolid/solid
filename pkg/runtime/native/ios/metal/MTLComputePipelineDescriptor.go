//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLComputePipelineDescriptor : objc.NSObject
*/

type MTLComputePipelineDescriptor struct {
  *objc.NSObject

}

func (r *MTLComputePipelineDescriptor) SetInsertLibraries() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetLinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SupportAddingBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) StageInputDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) InsertLibraries() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetBinaryArchives() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetSupportAddingBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) MaxCallStackDepth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) Label() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) ThreadGroupSizeIsMultipleOfThreadExecutionWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetSupportIndirectCommandBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetStageInputDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) PreloadedLibraries() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetPreloadedLibraries() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetMaxCallStackDepth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) Reset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) ComputeFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetThreadGroupSizeIsMultipleOfThreadExecutionWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) MaxTotalThreadsPerThreadgroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) Buffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) LinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetComputeFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetMaxTotalThreadsPerThreadgroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SupportIndirectCommandBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) BinaryArchives() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


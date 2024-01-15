//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLComputePipelineDescriptor : objc.NSObject
*/

type MTLComputePipelineDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLComputePipelineDescriptor) ThreadGroupSizeIsMultipleOfThreadExecutionWidth() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetBinaryArchives() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetThreadGroupSizeIsMultipleOfThreadExecutionWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) ComputeFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetComputeFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetInsertLibraries() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) LinkedFunctions() (*MTLLinkedFunctions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) MaxTotalThreadsPerThreadgroup() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetMaxTotalThreadsPerThreadgroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetPreloadedLibraries() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetLinkedFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) Reset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SupportAddingBinaryFunctions() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetSupportAddingBinaryFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetMaxCallStackDepth() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetStageInputDescriptor() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) Buffers() (*MTLPipelineBufferDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) InsertLibraries() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) PreloadedLibraries() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) StageInputDescriptor() (*MTLStageInputOutputDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) BinaryArchives() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SupportIndirectCommandBuffers() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) SetSupportIndirectCommandBuffers() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePipelineDescriptor) MaxCallStackDepth() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}


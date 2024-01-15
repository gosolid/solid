//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLMeshRenderPipelineDescriptor : objc.NSObject
*/

type MTLMeshRenderPipelineDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLMeshRenderPipelineDescriptor) SetRasterSampleCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetAlphaToOneEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MaxVertexAmplificationCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MaxTotalThreadsPerMeshThreadgroup() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetMaxTotalThreadgroupsPerMeshGrid() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MeshFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MaxTotalThreadsPerObjectThreadgroup() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetMaxTotalThreadsPerMeshThreadgroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) DepthAttachmentPixelFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) FragmentFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MeshBuffers() (*MTLPipelineBufferDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) RasterSampleCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetMaxVertexAmplificationCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) ColorAttachments() (*MTLRenderPipelineColorAttachmentDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) FragmentLinkedFunctions() (*MTLLinkedFunctions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetObjectFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetPayloadMemoryLength() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) IsAlphaToOneEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetRasterizationEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetStencilAttachmentPixelFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MeshLinkedFunctions() (*MTLLinkedFunctions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetMeshLinkedFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetDepthAttachmentPixelFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetFragmentFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MaxTotalThreadgroupsPerMeshGrid() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) IsAlphaToCoverageEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetAlphaToCoverageEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) IsRasterizationEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) ObjectLinkedFunctions() (*MTLLinkedFunctions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) Reset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetMaxTotalThreadsPerObjectThreadgroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) ObjectBuffers() (*MTLPipelineBufferDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) FragmentBuffers() (*MTLPipelineBufferDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) StencilAttachmentPixelFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetFragmentLinkedFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) ObjectFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetMeshFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) PayloadMemoryLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetObjectLinkedFunctions() error {
  return fmt.Errorf("unimplemented")
}


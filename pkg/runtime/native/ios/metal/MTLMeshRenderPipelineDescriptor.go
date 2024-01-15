//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Metal.MTLMeshRenderPipelineDescriptor : objc.NSObject
*/

type MTLMeshRenderPipelineDescriptor struct {
  *objc.NSObject

}

func (r *MTLMeshRenderPipelineDescriptor) SetObjectFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) FragmentLinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MaxTotalThreadgroupsPerMeshGrid() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetRasterizationEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) ObjectLinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SupportIndirectCommandBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetObjectLinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetRasterSampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) IsAlphaToCoverageEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetMaxTotalThreadgroupsPerMeshGrid() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) RasterSampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) IsRasterizationEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetDepthAttachmentPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MeshFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetMaxTotalThreadsPerMeshThreadgroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) ColorAttachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) StencilAttachmentPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetStencilAttachmentPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) Label() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) FragmentBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetMaxVertexAmplificationCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MeshBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MaxVertexAmplificationCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) DepthAttachmentPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetMeshLinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetFragmentLinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetMeshFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MaxTotalThreadsPerObjectThreadgroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MaxTotalThreadsPerMeshThreadgroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetPayloadMemoryLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) IsAlphaToOneEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) Reset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetFragmentFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetMaxTotalThreadsPerObjectThreadgroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) PayloadMemoryLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) ObjectBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetAlphaToCoverageEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetAlphaToOneEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetSupportIndirectCommandBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) ObjectFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) FragmentFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMeshRenderPipelineDescriptor) MeshLinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


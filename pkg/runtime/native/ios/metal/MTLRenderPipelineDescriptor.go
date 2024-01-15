//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLRenderPipelineDescriptor : objc.NSObject
*/

type MTLRenderPipelineDescriptor struct {
  *objc.NSObject

}

func (r *MTLRenderPipelineDescriptor) MaxFragmentCallStackDepth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetTessellationPartitionMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) IsTessellationFactorScaleEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SupportAddingVertexBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) TessellationOutputWindingOrder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetSupportIndirectCommandBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetVertexPreloadedLibraries() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) FragmentPreloadedLibraries() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) VertexLinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) FragmentFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) MaxVertexAmplificationCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) FragmentLinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) StencilAttachmentPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetTessellationFactorScaleEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) VertexBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) VertexFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetAlphaToCoverageEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetAlphaToOneEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetInputPrimitiveTopology() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetMaxFragmentCallStackDepth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) TessellationFactorFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetBinaryArchives() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetMaxVertexCallStackDepth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) IsRasterizationEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) TessellationPartitionMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) MaxTessellationFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetFragmentFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) DepthAttachmentPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetStencilAttachmentPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) IsAlphaToCoverageEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) VertexPreloadedLibraries() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SupportAddingFragmentBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) Label() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetTessellationFactorStepFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetSupportAddingFragmentBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) Reset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetDepthAttachmentPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SupportIndirectCommandBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) BinaryArchives() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetVertexFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) VertexDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetRasterSampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetVertexLinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) ColorAttachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetMaxTessellationFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetTessellationOutputWindingOrder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) IsAlphaToOneEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetFragmentLinkedFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) MaxVertexCallStackDepth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) TessellationFactorStepFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetTessellationFactorFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) FragmentBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetFragmentPreloadedLibraries() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetSupportAddingVertexBinaryFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetVertexDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetRasterizationEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetMaxVertexAmplificationCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) RasterSampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) InputPrimitiveTopology() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) TessellationControlPointIndexType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetSampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetTessellationControlPointIndexType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


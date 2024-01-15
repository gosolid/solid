//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLRenderPipelineDescriptor : objc.NSObject
*/

type MTLRenderPipelineDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLRenderPipelineDescriptor) BinaryArchives() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) FragmentPreloadedLibraries() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetSupportAddingFragmentBinaryFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) RasterSampleCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SupportIndirectCommandBuffers() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) TessellationFactorFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) MaxFragmentCallStackDepth() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) IsRasterizationEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) DepthAttachmentPixelFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetStencilAttachmentPixelFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) IsTessellationFactorScaleEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetTessellationFactorScaleEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetTessellationControlPointIndexType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) IsAlphaToCoverageEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) StencilAttachmentPixelFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetSupportIndirectCommandBuffers() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetVertexLinkedFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetVertexFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetTessellationOutputWindingOrder() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SupportAddingVertexBinaryFunctions() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetSupportAddingVertexBinaryFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetMaxFragmentCallStackDepth() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) Reset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetDepthAttachmentPixelFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetInputPrimitiveTopology() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetTessellationFactorStepFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) TessellationOutputWindingOrder() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetAlphaToCoverageEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetMaxVertexAmplificationCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) MaxTessellationFactor() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetBinaryArchives() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetVertexPreloadedLibraries() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetAlphaToOneEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) TessellationPartitionMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) TessellationControlPointIndexType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) TessellationFactorStepFunction() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) FragmentLinkedFunctions() (*MTLLinkedFunctions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) VertexDescriptor() (*MTLVertexDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetVertexDescriptor() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) VertexBuffers() (*MTLPipelineBufferDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetFragmentLinkedFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) MaxVertexCallStackDepth() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) VertexFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetSampleCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetTessellationFactorFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetFragmentPreloadedLibraries() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) InputPrimitiveTopology() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetMaxTessellationFactor() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) IsAlphaToOneEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetTessellationPartitionMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) MaxVertexAmplificationCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) VertexPreloadedLibraries() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SupportAddingFragmentBinaryFunctions() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) FragmentFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SampleCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) VertexLinkedFunctions() (*MTLLinkedFunctions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetRasterSampleCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) FragmentBuffers() (*MTLPipelineBufferDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetMaxVertexCallStackDepth() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) ColorAttachments() (*MTLRenderPipelineColorAttachmentDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetFragmentFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineDescriptor) SetRasterizationEnabled() error {
  return fmt.Errorf("unimplemented")
}


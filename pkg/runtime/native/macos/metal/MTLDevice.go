//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLDevice
*/

type MTLDevice struct {

}

func (r *MTLDevice) NewHeapWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsTextureSampleCount() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewBinaryArchiveWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) CurrentAllocatedSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SparseTileSizeInBytes() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsRenderDynamicLibraries() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewBufferWithBytes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsCounterSampling() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) MaxThreadgroupMemoryLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) MaxArgumentBufferSamplerCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) HeapBufferSizeAndAlignWithLength() (MTLSizeAndAlign, error) {
  return MTLSizeAndAlign{}, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewRenderPipelineStateWithMeshDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) Supports32BitMSAA() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) MaxBufferLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsRaytracing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewBufferWithLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewTextureWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsFeatureSet() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) Architecture() (*MTLArchitecture, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsQueryTextureLOD() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) GetDefaultSamplePositions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewDynamicLibraryWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewAccelerationStructureWithSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) HasUnifiedMemory() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewSharedTextureWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsFamily() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewRenderPipelineStateWithTileDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewRasterizationRateMapWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SparseTileSizeWithTextureType() (MTLSize, error) {
  return MTLSize{}, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SparseTileSizeInBytesForSparsePageSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewSamplerStateWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewLibraryWithStitchedDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewComputePipelineStateWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) MinimumLinearTextureAlignmentForPixelFormat() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) AccelerationStructureSizesWithDescriptor() (MTLAccelerationStructureSizes, error) {
  return MTLAccelerationStructureSizes{}, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) ArgumentBuffersSupport() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) PeerCount() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) MaximumConcurrentCompilationTaskCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewLibraryWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewCounterSampleBufferWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SampleTimestamps() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsRasterizationRateMapWithLayerCount() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) Location() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewLibraryWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) CounterSets() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) ShouldMaximizeConcurrentCompilation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewSharedTextureWithHandle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) MaxThreadsPerThreadgroup() (MTLSize, error) {
  return MTLSize{}, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) ReadWriteTextureSupport() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) IsDepth24Stencil8PixelFormatSupported() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsFunctionPointersFromRender() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewDepthStencilStateWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewIOCommandQueueWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsRaytracingFromRender() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewLibraryWithSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) LocationNumber() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsPullModelInterpolation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) PeerIndex() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsFunctionPointers() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewIOFileHandleWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsVertexAmplificationCount() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) ConvertSparsePixelRegions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsBCTextureCompression() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsDynamicLibraries() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsPrimitiveMotionBlur() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewCommandQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewAccelerationStructureWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) RegistryID() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) RecommendedMaxWorkingSetSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SetShouldMaximizeConcurrentCompilation() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewIOHandleWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) IsRemovable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewDefaultLibraryWithBundle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewComputePipelineStateWithFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) AreRasterOrderGroupsSupported() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewBufferWithBytesNoCopy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) MinimumTextureBufferAlignmentForPixelFormat() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewSharedEventWithHandle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) HeapAccelerationStructureSizeAndAlignWithDescriptor() (MTLSizeAndAlign, error) {
  return MTLSizeAndAlign{}, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) MaxTransferRate() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) SupportsShaderBarycentricCoordinates() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewLibraryWithFile() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewFence() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewIndirectCommandBufferWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewSharedEvent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) AreBarycentricCoordsSupported() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewArgumentEncoderWithArguments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) ConvertSparseTileRegions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewArgumentEncoderWithBufferBinding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) IsLowPower() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) HeapTextureSizeAndAlignWithDescriptor() (MTLSizeAndAlign, error) {
  return MTLSizeAndAlign{}, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) HeapAccelerationStructureSizeAndAlignWithSize() (MTLSizeAndAlign, error) {
  return MTLSizeAndAlign{}, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) IsHeadless() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) AreProgrammableSamplePositionsSupported() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) PeerGroupID() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewRenderPipelineStateWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewCommandQueueWithMaxCommandBufferCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewDefaultLibrary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewDynamicLibrary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) NewEvent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDevice) Supports32BitFloatFiltering() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLTextureDescriptor : objc.NSObject
*/

type MTLTextureDescriptor struct {
  *objc.NSObject

}

func (r *MTLTextureDescriptor) Texture2DDescriptorWithPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) AllowGPUOptimizedContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetSwizzle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) HazardTrackingMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) Width() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetResourceOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) StorageMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetCpuCacheMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetHazardTrackingMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetUsage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetTextureType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetMipmapLevelCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) CpuCacheMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetCompressionType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) TextureCubeDescriptorWithPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetSampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetArrayLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) ArrayLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetStorageMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) Swizzle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) TextureType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) MipmapLevelCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) Usage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) PixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) Depth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) TextureBufferDescriptorWithPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetDepth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) CompressionType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) Height() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) ResourceOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetAllowGPUOptimizedContents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


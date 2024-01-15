//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLTextureDescriptor : objc.NSObject
*/

type MTLTextureDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLTextureDescriptor) SetResourceOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetCpuCacheMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) StorageMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetStorageMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetAllowGPUOptimizedContents() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) CompressionType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) TextureBufferDescriptorWithPixelFormat() (*MTLTextureDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) TextureType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetDepth() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) Texture2DDescriptorWithPixelFormat() (*MTLTextureDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetHazardTrackingMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) HazardTrackingMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetSwizzle() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetTextureType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) Width() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) Depth() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetSampleCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetArrayLength() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) ResourceOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) Usage() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetCompressionType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) TextureCubeDescriptorWithPixelFormat() (*MTLTextureDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) PixelFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) MipmapLevelCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetUsage() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) Height() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SampleCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) ArrayLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) AllowGPUOptimizedContents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) Swizzle() (MTLTextureSwizzleChannels, error) {
  return MTLTextureSwizzleChannels{}, fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetPixelFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) SetMipmapLevelCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTextureDescriptor) CpuCacheMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


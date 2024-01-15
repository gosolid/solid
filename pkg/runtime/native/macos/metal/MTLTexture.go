//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
)

/*
protocol Metal.MTLTexture
*/

type MTLTexture struct {

}

func (r *MTLTexture) ParentTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) BufferBytesPerRow() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) Swizzle() (MTLTextureSwizzleChannels, error) {
  return MTLTextureSwizzleChannels{}, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) Width() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) GetBytes() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTexture) ReplaceRegion() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLTexture) NewRemoteTextureViewForDevice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) MipmapLevelCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) IsFramebufferOnly() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) NewSharedTextureHandle() (*MTLSharedTextureHandle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) Iosurface() (*core_graphics.IOSurface, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) Height() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) ArrayLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) CompressionType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) BufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) Depth() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) IsSparse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) GpuResourceID() (MTLResourceID, error) {
  return MTLResourceID{}, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) RemoteStorageTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) ParentRelativeLevel() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) ParentRelativeSlice() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) Buffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) PixelFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) SampleCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) FirstMipmapInTail() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) TailSizeInBytes() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) NewTextureViewWithPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) RootResource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) IosurfacePlane() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) TextureType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) IsShareable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) AllowGPUOptimizedContents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLTexture) Usage() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_video"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
)

/*
protocol CoreImage.CIImageProcessorInput
*/

type CIImageProcessorInput struct {

}

func (r *CIImageProcessorInput) BaseAddress() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorInput) PixelBuffer() (*core_video.CVBuffer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorInput) MetalTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorInput) Region() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorInput) BytesPerRow() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorInput) Format() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorInput) Surface() (*core_graphics.IOSurface, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorInput) Digest() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorInput) RoiTileIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorInput) RoiTileCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}


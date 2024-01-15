//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_video"
)

/*
protocol CoreImage.CIImageProcessorOutput
*/

type CIImageProcessorOutput struct {

}

func (r *CIImageProcessorOutput) Format() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorOutput) BaseAddress() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorOutput) Surface() (*core_graphics.IOSurface, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorOutput) Digest() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorOutput) Region() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorOutput) BytesPerRow() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorOutput) PixelBuffer() (*core_video.CVBuffer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorOutput) MetalTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorOutput) MetalCommandBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
)

/*
interface CoreImage.CIRenderDestination : objc.NSObject
*/

type CIRenderDestination struct {
  *objc.NSObject

}

func (r *CIRenderDestination) InitWithMTLTexture() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) Height() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) ColorSpace() (*core_graphics.CGColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) BlendsInDestinationColorSpace() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) SetBlendsInDestinationColorSpace() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) InitWithBitmapData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) SetAlphaMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) SetFlipped() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) BlendKernel() (*CIBlendKernel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) SetColorSpace() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) InitWithPixelBuffer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) InitWithIOSurface() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) Width() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) AlphaMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) IsFlipped() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) IsDithered() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) InitWithWidth() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) InitWithGLTexture() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) SetDithered() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) IsClamped() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) SetClamped() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) SetBlendKernel() error {
  return fmt.Errorf("unimplemented")
}


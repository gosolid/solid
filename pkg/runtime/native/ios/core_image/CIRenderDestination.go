//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CIRenderDestination : objc.NSObject
*/

type CIRenderDestination struct {
  *objc.NSObject

}

func (r *CIRenderDestination) Height() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) AlphaMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) SetFlipped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) IsClamped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) ColorSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) SetColorSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) SetBlendKernel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) InitWithGLTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) InitWithIOSurface() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) Width() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) SetDithered() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) BlendsInDestinationColorSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) InitWithPixelBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) InitWithWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) InitWithBitmapData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) IsDithered() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) BlendKernel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) InitWithMTLTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) IsFlipped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) SetClamped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) SetBlendsInDestinationColorSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderDestination) SetAlphaMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


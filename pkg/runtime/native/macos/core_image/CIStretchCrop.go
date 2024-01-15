//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIStretchCrop
*/

type CIStretchCrop struct {

}

func (r *CIStretchCrop) SetCropAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStretchCrop) CenterStretchAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIStretchCrop) SetCenterStretchAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStretchCrop) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIStretchCrop) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStretchCrop) Size() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIStretchCrop) SetSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStretchCrop) CropAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}


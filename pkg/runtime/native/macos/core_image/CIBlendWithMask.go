//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIBlendWithMask
*/

type CIBlendWithMask struct {

}

func (r *CIBlendWithMask) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIBlendWithMask) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBlendWithMask) BackgroundImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIBlendWithMask) SetBackgroundImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBlendWithMask) MaskImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIBlendWithMask) SetMaskImage() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIUnsharpMask
*/

type CIUnsharpMask struct {

}

func (r *CIUnsharpMask) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIUnsharpMask) Intensity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIUnsharpMask) SetIntensity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIUnsharpMask) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIUnsharpMask) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIUnsharpMask) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIHeightFieldFromMask
*/

type CIHeightFieldFromMask struct {

}

func (r *CIHeightFieldFromMask) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIHeightFieldFromMask) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHeightFieldFromMask) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIHeightFieldFromMask) SetRadius() error {
  return fmt.Errorf("unimplemented")
}


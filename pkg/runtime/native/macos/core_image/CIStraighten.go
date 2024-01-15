//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIStraighten
*/

type CIStraighten struct {

}

func (r *CIStraighten) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIStraighten) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStraighten) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIStraighten) SetAngle() error {
  return fmt.Errorf("unimplemented")
}


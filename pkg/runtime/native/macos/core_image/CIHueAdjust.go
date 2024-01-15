//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIHueAdjust
*/

type CIHueAdjust struct {

}

func (r *CIHueAdjust) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIHueAdjust) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHueAdjust) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIHueAdjust) SetAngle() error {
  return fmt.Errorf("unimplemented")
}


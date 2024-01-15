//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIMotionBlur
*/

type CIMotionBlur struct {

}

func (r *CIMotionBlur) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMotionBlur) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMotionBlur) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMotionBlur) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIMotionBlur) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMotionBlur) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}


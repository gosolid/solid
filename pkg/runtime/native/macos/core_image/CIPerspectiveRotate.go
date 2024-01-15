//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIPerspectiveRotate
*/

type CIPerspectiveRotate struct {

}

func (r *CIPerspectiveRotate) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveRotate) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveRotate) FocalLength() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveRotate) Yaw() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveRotate) SetYaw() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveRotate) Roll() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveRotate) SetFocalLength() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveRotate) Pitch() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveRotate) SetPitch() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveRotate) SetRoll() error {
  return fmt.Errorf("unimplemented")
}


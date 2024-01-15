//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIDisplacementDistortion
*/

type CIDisplacementDistortion struct {

}

func (r *CIDisplacementDistortion) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDisplacementDistortion) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDisplacementDistortion) DisplacementImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDisplacementDistortion) SetDisplacementImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDisplacementDistortion) Scale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDisplacementDistortion) SetScale() error {
  return fmt.Errorf("unimplemented")
}


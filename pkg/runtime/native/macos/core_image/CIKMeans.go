//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIKMeans
*/

type CIKMeans struct {

}

func (r *CIKMeans) SetPasses() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIKMeans) Perceptual() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKMeans) SetPerceptual() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIKMeans) InputMeans() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKMeans) SetInputMeans() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIKMeans) Count() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIKMeans) SetCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIKMeans) Passes() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}


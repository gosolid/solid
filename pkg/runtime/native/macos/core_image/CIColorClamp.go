//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIColorClamp
*/

type CIColorClamp struct {

}

func (r *CIColorClamp) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorClamp) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorClamp) MinComponents() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorClamp) SetMinComponents() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorClamp) MaxComponents() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorClamp) SetMaxComponents() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIConvertLab
*/

type CIConvertLab struct {

}

func (r *CIConvertLab) Normalize() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIConvertLab) SetNormalize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIConvertLab) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIConvertLab) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}


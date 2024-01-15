//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIPersonSegmentation
*/

type CIPersonSegmentation struct {

}

func (r *CIPersonSegmentation) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPersonSegmentation) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPersonSegmentation) QualityLevel() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPersonSegmentation) SetQualityLevel() error {
  return fmt.Errorf("unimplemented")
}


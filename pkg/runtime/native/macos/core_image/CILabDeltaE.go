//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CILabDeltaE
*/

type CILabDeltaE struct {

}

func (r *CILabDeltaE) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CILabDeltaE) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILabDeltaE) Image2() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CILabDeltaE) SetImage2() error {
  return fmt.Errorf("unimplemented")
}


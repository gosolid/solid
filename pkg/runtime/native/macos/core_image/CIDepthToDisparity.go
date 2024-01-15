//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIDepthToDisparity
*/

type CIDepthToDisparity struct {

}

func (r *CIDepthToDisparity) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDepthToDisparity) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}


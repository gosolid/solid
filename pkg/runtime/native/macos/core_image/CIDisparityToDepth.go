//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIDisparityToDepth
*/

type CIDisparityToDepth struct {

}

func (r *CIDisparityToDepth) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDisparityToDepth) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}


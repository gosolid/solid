//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIMedian
*/

type CIMedian struct {

}

func (r *CIMedian) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMedian) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}


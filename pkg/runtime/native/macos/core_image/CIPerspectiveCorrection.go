//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIPerspectiveCorrection
*/

type CIPerspectiveCorrection struct {

}

func (r *CIPerspectiveCorrection) Crop() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveCorrection) SetCrop() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIAffineClamp
*/

type CIAffineClamp struct {

}

func (r *CIAffineClamp) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIAffineClamp) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAffineClamp) Transform() (core_foundation.CGAffineTransform, error) {
  return core_foundation.CGAffineTransform{}, fmt.Errorf("unimplemented")
}

func (r *CIAffineClamp) SetTransform() error {
  return fmt.Errorf("unimplemented")
}


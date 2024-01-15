//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIAffineTile
*/

type CIAffineTile struct {

}

func (r *CIAffineTile) SetTransform() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAffineTile) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIAffineTile) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAffineTile) Transform() (core_foundation.CGAffineTransform, error) {
  return core_foundation.CGAffineTransform{}, fmt.Errorf("unimplemented")
}


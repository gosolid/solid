//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIZoomBlur
*/

type CIZoomBlur struct {

}

func (r *CIZoomBlur) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIZoomBlur) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIZoomBlur) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIZoomBlur) Amount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIZoomBlur) SetAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIZoomBlur) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}


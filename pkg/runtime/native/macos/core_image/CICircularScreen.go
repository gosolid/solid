//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
)

/*
protocol CoreImage.CICircularScreen
*/

type CICircularScreen struct {

}

func (r *CICircularScreen) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICircularScreen) Sharpness() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICircularScreen) SetSharpness() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICircularScreen) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICircularScreen) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICircularScreen) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CICircularScreen) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICircularScreen) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}


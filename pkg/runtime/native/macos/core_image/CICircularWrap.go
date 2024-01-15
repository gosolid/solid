//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CICircularWrap
*/

type CICircularWrap struct {

}

func (r *CICircularWrap) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICircularWrap) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICircularWrap) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICircularWrap) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICircularWrap) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICircularWrap) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CICircularWrap) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICircularWrap) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}


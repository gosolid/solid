//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CISwipeTransition
*/

type CISwipeTransition struct {

}

func (r *CISwipeTransition) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISwipeTransition) Color() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISwipeTransition) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISwipeTransition) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISwipeTransition) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISwipeTransition) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISwipeTransition) Opacity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISwipeTransition) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CISwipeTransition) SetExtent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISwipeTransition) SetOpacity() error {
  return fmt.Errorf("unimplemented")
}


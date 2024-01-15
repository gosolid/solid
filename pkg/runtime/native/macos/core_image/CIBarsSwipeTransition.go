//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIBarsSwipeTransition
*/

type CIBarsSwipeTransition struct {

}

func (r *CIBarsSwipeTransition) SetBarOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBarsSwipeTransition) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBarsSwipeTransition) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBarsSwipeTransition) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBarsSwipeTransition) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBarsSwipeTransition) BarOffset() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}


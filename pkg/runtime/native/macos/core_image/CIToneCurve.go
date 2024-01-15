//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIToneCurve
*/

type CIToneCurve struct {

}

func (r *CIToneCurve) Point1() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIToneCurve) Point2() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIToneCurve) Point3() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIToneCurve) Point4() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIToneCurve) SetPoint4() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIToneCurve) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIToneCurve) Point0() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIToneCurve) SetPoint0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIToneCurve) SetPoint3() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIToneCurve) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIToneCurve) SetPoint1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIToneCurve) SetPoint2() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CICopyMachineTransition
*/

type CICopyMachineTransition struct {

}

func (r *CICopyMachineTransition) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICopyMachineTransition) Opacity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICopyMachineTransition) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CICopyMachineTransition) Color() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICopyMachineTransition) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICopyMachineTransition) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICopyMachineTransition) SetExtent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICopyMachineTransition) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICopyMachineTransition) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICopyMachineTransition) SetOpacity() error {
  return fmt.Errorf("unimplemented")
}


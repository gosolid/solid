//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CINinePartStretched
*/

type CINinePartStretched struct {

}

func (r *CINinePartStretched) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CINinePartStretched) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CINinePartStretched) Breakpoint0() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CINinePartStretched) SetBreakpoint0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CINinePartStretched) Breakpoint1() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CINinePartStretched) SetBreakpoint1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CINinePartStretched) GrowAmount() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CINinePartStretched) SetGrowAmount() error {
  return fmt.Errorf("unimplemented")
}


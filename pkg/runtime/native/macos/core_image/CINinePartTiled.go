//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CINinePartTiled
*/

type CINinePartTiled struct {

}

func (r *CINinePartTiled) Breakpoint1() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CINinePartTiled) SetBreakpoint1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CINinePartTiled) FlipYTiles() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CINinePartTiled) SetFlipYTiles() error {
  return fmt.Errorf("unimplemented")
}

func (r *CINinePartTiled) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CINinePartTiled) SetBreakpoint0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CINinePartTiled) GrowAmount() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CINinePartTiled) SetGrowAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CINinePartTiled) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CINinePartTiled) Breakpoint0() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}


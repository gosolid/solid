//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CICheckerboardGenerator
*/

type CICheckerboardGenerator struct {

}

func (r *CICheckerboardGenerator) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICheckerboardGenerator) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CICheckerboardGenerator) Color0() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICheckerboardGenerator) SetColor0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICheckerboardGenerator) Color1() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICheckerboardGenerator) SetSharpness() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICheckerboardGenerator) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICheckerboardGenerator) SetColor1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICheckerboardGenerator) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICheckerboardGenerator) Sharpness() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}


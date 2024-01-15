//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIVignetteEffect
*/

type CIVignetteEffect struct {

}

func (r *CIVignetteEffect) SetIntensity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIVignetteEffect) SetFalloff() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIVignetteEffect) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVignetteEffect) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIVignetteEffect) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIVignetteEffect) Intensity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIVignetteEffect) Falloff() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIVignetteEffect) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIVignetteEffect) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIVignetteEffect) SetRadius() error {
  return fmt.Errorf("unimplemented")
}


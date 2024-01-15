//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CICircleSplashDistortion
*/

type CICircleSplashDistortion struct {

}

func (r *CICircleSplashDistortion) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CICircleSplashDistortion) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICircleSplashDistortion) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICircleSplashDistortion) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICircleSplashDistortion) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICircleSplashDistortion) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}


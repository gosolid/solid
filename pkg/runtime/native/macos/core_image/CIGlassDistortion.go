//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIGlassDistortion
*/

type CIGlassDistortion struct {

}

func (r *CIGlassDistortion) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIGlassDistortion) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGlassDistortion) TextureImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIGlassDistortion) SetTextureImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGlassDistortion) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIGlassDistortion) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGlassDistortion) Scale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIGlassDistortion) SetScale() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CILightTunnel
*/

type CILightTunnel struct {

}

func (r *CILightTunnel) SetRotation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILightTunnel) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILightTunnel) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILightTunnel) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CILightTunnel) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILightTunnel) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CILightTunnel) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILightTunnel) Rotation() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
)

/*
protocol CoreImage.CIParallelogramTile
*/

type CIParallelogramTile struct {

}

func (r *CIParallelogramTile) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIParallelogramTile) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIParallelogramTile) AcuteAngle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIParallelogramTile) SetAcuteAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIParallelogramTile) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIParallelogramTile) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIParallelogramTile) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIParallelogramTile) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIParallelogramTile) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIParallelogramTile) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIPointillize
*/

type CIPointillize struct {

}

func (r *CIPointillize) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPointillize) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPointillize) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPointillize) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPointillize) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIPointillize) SetCenter() error {
  return fmt.Errorf("unimplemented")
}


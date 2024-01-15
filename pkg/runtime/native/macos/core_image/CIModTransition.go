//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIModTransition
*/

type CIModTransition struct {

}

func (r *CIModTransition) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIModTransition) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIModTransition) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIModTransition) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIModTransition) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIModTransition) Compression() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIModTransition) SetCompression() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIModTransition) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}


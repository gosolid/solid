//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSRotationGestureRecognizer : AppKit.NSGestureRecognizer
*/

type NSRotationGestureRecognizer struct {
  *NSGestureRecognizer

}

func (r *NSRotationGestureRecognizer) Rotation() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRotationGestureRecognizer) SetRotation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRotationGestureRecognizer) RotationInDegrees() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRotationGestureRecognizer) SetRotationInDegrees() error {
  return fmt.Errorf("unimplemented")
}


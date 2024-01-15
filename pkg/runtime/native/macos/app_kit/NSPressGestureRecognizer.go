//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSPressGestureRecognizer : AppKit.NSGestureRecognizer
*/

type NSPressGestureRecognizer struct {
  *NSGestureRecognizer
  *foundation.NSCoding
}

func (r *NSPressGestureRecognizer) SetAllowableMovement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPressGestureRecognizer) NumberOfTouchesRequired() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPressGestureRecognizer) SetNumberOfTouchesRequired() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPressGestureRecognizer) ButtonMask() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPressGestureRecognizer) SetButtonMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPressGestureRecognizer) MinimumPressDuration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPressGestureRecognizer) SetMinimumPressDuration() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPressGestureRecognizer) AllowableMovement() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}


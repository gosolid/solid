//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSPanGestureRecognizer : AppKit.NSGestureRecognizer
*/

type NSPanGestureRecognizer struct {
  *NSGestureRecognizer
  *foundation.NSCoding
}

func (r *NSPanGestureRecognizer) SetTranslation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPanGestureRecognizer) VelocityInView() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSPanGestureRecognizer) ButtonMask() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPanGestureRecognizer) SetButtonMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPanGestureRecognizer) NumberOfTouchesRequired() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPanGestureRecognizer) SetNumberOfTouchesRequired() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPanGestureRecognizer) TranslationInView() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}


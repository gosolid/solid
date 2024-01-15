//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSGestureRecognizer : objc.NSObject
*/

type NSGestureRecognizer struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSGestureRecognizer) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) DelaysMagnificationEvents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) SetDelaysRotationEvents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) Target() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) InitWithTarget() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) SetAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) SetDelaysOtherMouseButtonEvents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) Action() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) View() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) DelaysPrimaryMouseButtonEvents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) SetPressureConfiguration() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) DelaysRotationEvents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) SetDelaysMagnificationEvents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) SetTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) State() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) IsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) PressureConfiguration() (*NSPressureConfiguration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) SetDelaysPrimaryMouseButtonEvents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) DelaysSecondaryMouseButtonEvents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) SetDelaysSecondaryMouseButtonEvents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) DelaysKeyEvents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) LocationInView() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) DelaysOtherMouseButtonEvents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGestureRecognizer) SetDelaysKeyEvents() error {
  return fmt.Errorf("unimplemented")
}


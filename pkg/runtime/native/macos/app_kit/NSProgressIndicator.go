//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSProgressIndicator : AppKit.NSView
*/

type NSProgressIndicator struct {
  *NSView
  *NSAccessibilityProgressIndicator
}

func (r *NSProgressIndicator) SetMinValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) SetMaxValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) SetUsesThreadedAnimation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) Style() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) StartAnimation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) SizeToFit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) ControlSize() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) MinValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) SetStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) UsesThreadedAnimation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) IncrementBy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) SetControlSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) MaxValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) SetObservedProgress() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) IsIndeterminate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) SetDoubleValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) IsDisplayedWhenStopped() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) SetDisplayedWhenStopped() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) StopAnimation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) SetIndeterminate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) DoubleValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSProgressIndicator) ObservedProgress() (*foundation.NSProgress, error) {
  return nil, fmt.Errorf("unimplemented")
}


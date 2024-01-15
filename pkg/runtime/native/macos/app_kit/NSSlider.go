//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSSlider : AppKit.NSControl
*/

type NSSlider struct {
  *NSControl
  *NSAccessibilitySlider
}

func (r *NSSlider) SetMaxValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSlider) IsVertical() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSlider) SetVertical() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSlider) TrackFillColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSlider) SetAltIncrementValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSlider) KnobThickness() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSlider) AcceptsFirstMouse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSlider) SliderType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSlider) SetSliderType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSlider) MaxValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSlider) MinValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSlider) SetMinValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSlider) AltIncrementValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSlider) SetTrackFillColor() error {
  return fmt.Errorf("unimplemented")
}


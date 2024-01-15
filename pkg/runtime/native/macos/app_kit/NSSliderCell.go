//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSSliderCell : AppKit.NSActionCell
*/

type NSSliderCell struct {
  *NSActionCell

}

func (r *NSSliderCell) BarRectFlipped() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) PrefersTrackingUntilMouseUp() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) MaxValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) SetVertical() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) DrawBarInside() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) MinValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) SetMaxValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) SetAltIncrementValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) IsVertical() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) DrawKnob() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) AltIncrementValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) TrackRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) KnobThickness() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) KnobRectFlipped() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) SetMinValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) SliderType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSliderCell) SetSliderType() error {
  return fmt.Errorf("unimplemented")
}


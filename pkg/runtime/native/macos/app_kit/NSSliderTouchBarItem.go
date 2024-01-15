//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSSliderTouchBarItem : AppKit.NSTouchBarItem
*/

type NSSliderTouchBarItem struct {
  *NSTouchBarItem

}

func (r *NSSliderTouchBarItem) SetAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) SetMinimumSliderWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) MinimumValueAccessory() (*NSSliderAccessory, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) SetMinimumValueAccessory() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) Action() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) MaximumSliderWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) SetCustomizationLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) View() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) Slider() (*NSSlider, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) DoubleValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) SetDoubleValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) MaximumValueAccessory() (*NSSliderAccessory, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) SetValueAccessoryWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) Target() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) CustomizationLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) SetSlider() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) MinimumSliderWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) SetMaximumSliderWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) SetMaximumValueAccessory() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) ValueAccessoryWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSliderTouchBarItem) SetTarget() error {
  return fmt.Errorf("unimplemented")
}


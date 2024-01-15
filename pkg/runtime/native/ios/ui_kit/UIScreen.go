//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIScreen : objc.NSObject
*/

type UIScreen struct {
  *objc.NSObject

}

func (r *UIScreen) OverscanCompensation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) MirroredScreen() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) CoordinateSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) CurrentEDRHeadroom() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) Screens() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) WantsSoftwareDimming() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) AvailableModes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) FixedCoordinateSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) NativeScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) MaximumFramesPerSecond() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) FocusedItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) MainScreen() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) Bounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) Scale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) CurrentMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) OverscanCompensationInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) FocusedView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) ApplicationFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) NativeBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) CalibratedLatency() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) PotentialEDRHeadroom() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) DisplayLinkWithTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) SetCurrentMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) SetOverscanCompensation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) SetWantsSoftwareDimming() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) Brightness() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) ReferenceDisplayModeStatus() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) SupportsFocus() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) PreferredMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) IsCaptured() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreen) SetBrightness() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSScroller : AppKit.NSControl
*/

type NSScroller struct {
  *NSControl

}

func (r *NSScroller) DrawKnob() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScroller) TrackKnob() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScroller) SetKnobProportion() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScroller) UsableParts() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScroller) KnobProportion() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScroller) PreferredScrollerStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScroller) ScrollerStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScroller) SetKnobStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScroller) IsCompatibleWithOverlayScrollers() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScroller) ScrollerWidthForControlSize() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScroller) CheckSpaceForParts() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScroller) TestPart() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScroller) KnobStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScroller) ControlSize() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScroller) SetControlSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScroller) HitPart() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScroller) RectForPart() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSScroller) DrawKnobSlotInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScroller) SetScrollerStyle() error {
  return fmt.Errorf("unimplemented")
}


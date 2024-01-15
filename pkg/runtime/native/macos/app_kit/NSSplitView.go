//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSSplitView : AppKit.NSView
*/

type NSSplitView struct {
  *NSView

}

func (r *NSSplitView) MaxPossiblePositionOfDividerAtIndex() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSplitView) HoldingPriorityForSubviewAtIndex() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSplitView) IsVertical() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSplitView) AutosaveName() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSplitView) DrawDividerInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitView) SetPosition() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitView) SetHoldingPriority() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitView) DividerStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSplitView) SetDividerStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitView) DividerColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSplitView) IsSubviewCollapsed() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSplitView) MinPossiblePositionOfDividerAtIndex() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSplitView) SetVertical() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitView) DividerThickness() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSplitView) AdjustSubviews() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitView) SetAutosaveName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitView) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSplitView) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}


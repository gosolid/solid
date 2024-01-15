//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSStackView : AppKit.NSView
*/

type NSStackView struct {
  *NSView

}

func (r *NSStackView) ArrangedSubviews() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStackView) AddArrangedSubview() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStackView) DetachesHiddenViews() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSStackView) SetDistribution() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStackView) Spacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStackView) RemoveArrangedSubview() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStackView) SetVisibilityPriority() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStackView) ClippingResistancePriorityForOrientation() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStackView) EdgeInsets() (foundation.NSEdgeInsets, error) {
  return foundation.NSEdgeInsets{}, fmt.Errorf("unimplemented")
}

func (r *NSStackView) SetCustomSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStackView) CustomSpacingAfterView() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStackView) SetHuggingPriority() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStackView) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStackView) VisibilityPriorityForView() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStackView) HuggingPriorityForOrientation() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStackView) Orientation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStackView) SetOrientation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStackView) SetAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStackView) SetSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStackView) StackViewWithViews() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStackView) InsertArrangedSubview() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStackView) SetDetachesHiddenViews() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStackView) Distribution() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStackView) DetachedViews() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStackView) SetClippingResistancePriority() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStackView) Alignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStackView) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStackView) SetEdgeInsets() error {
  return fmt.Errorf("unimplemented")
}


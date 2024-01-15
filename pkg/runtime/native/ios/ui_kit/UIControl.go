//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIControl : UIKit.UIView
*/

type UIControl struct {
  *UIView

}

func (r *UIControl) SendActionsForControlEvents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) SetSelected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) SetHighlighted() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) ToolTipInteraction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) SetContentHorizontalAlignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) IsTouchInside() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) AllTargets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) SetToolTip() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) ToolTip() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) SetContentVerticalAlignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) ContentHorizontalAlignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) SetContextMenuInteractionEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) ShowsMenuAsPrimaryAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) AddAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) RemoveAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) SetEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) State() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) ContextMenuInteraction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) IsContextMenuInteractionEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) IsSymbolAnimationEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) RemoveActionForIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) ActionsForTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) EnumerateEventHandlers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) ContentVerticalAlignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) SetSymbolAnimationEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) IsTracking() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) SetShowsMenuAsPrimaryAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) BeginTrackingWithTouch() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) CancelTrackingWithEvent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) AddTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) SendAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) AllControlEvents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) InitWithFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) RemoveTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) IsSelected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) IsHighlighted() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) EffectiveContentHorizontalAlignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) ContinueTrackingWithTouch() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) EndTrackingWithTouch() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIControl) MenuAttachmentPointForConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


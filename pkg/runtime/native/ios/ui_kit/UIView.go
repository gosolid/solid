//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIView : UIKit.UIResponder
*/

type UIView struct {
  *UIResponder

}

func (r *UIView) SetUserInteractionEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) SetTag() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) SetSemanticContentAttribute() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) EffectiveUserInterfaceLayoutDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) FocusGroupPriority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) SetFocusGroupPriority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) InitWithFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) LayerClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) IsUserInteractionEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) Layer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) CanBecomeFocused() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) SetFocusGroupIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) SetFocusEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) FocusEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) SemanticContentAttribute() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) UserInterfaceLayoutDirectionForSemanticContentAttribute() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) Tag() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) IsFocused() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIView) FocusGroupIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


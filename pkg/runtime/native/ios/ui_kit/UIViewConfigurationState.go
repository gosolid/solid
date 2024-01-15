//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIViewConfigurationState : objc.NSObject
*/

type UIViewConfigurationState struct {
  *objc.NSObject

}

func (r *UIViewConfigurationState) SetDisabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) IsFocused() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) TraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) IsDisabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) IsHighlighted() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) SetSelected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) IsPinned() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) SetPinned() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) SetTraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) IsSelected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) SetHighlighted() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) SetFocused() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewConfigurationState) InitWithTraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


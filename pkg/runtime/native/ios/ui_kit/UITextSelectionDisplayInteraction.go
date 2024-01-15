//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UITextSelectionDisplayInteraction : objc.NSObject
*/

type UITextSelectionDisplayInteraction struct {
  *objc.NSObject

}

func (r *UITextSelectionDisplayInteraction) HighlightView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionDisplayInteraction) SetNeedsSelectionUpdate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionDisplayInteraction) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionDisplayInteraction) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionDisplayInteraction) SetCursorView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionDisplayInteraction) SetHandleViews() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionDisplayInteraction) LayoutManagedSubviews() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionDisplayInteraction) IsActivated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionDisplayInteraction) SetActivated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionDisplayInteraction) HandleViews() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionDisplayInteraction) TextInput() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionDisplayInteraction) CursorView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionDisplayInteraction) SetHighlightView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionDisplayInteraction) InitWithTextInput() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionDisplayInteraction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIAccessibilityCustomAction : objc.NSObject
*/

type UIAccessibilityCustomAction struct {
  *objc.NSObject

}

func (r *UIAccessibilityCustomAction) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomAction) SetAttributedName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomAction) SetTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomAction) SetSelector() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomAction) InitWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomAction) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomAction) Selector() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomAction) SetActionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomAction) InitWithAttributedName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomAction) SetImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomAction) Image() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomAction) ActionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomAction) AttributedName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomAction) Target() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


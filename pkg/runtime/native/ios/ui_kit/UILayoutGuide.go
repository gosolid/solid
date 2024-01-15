//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UILayoutGuide : objc.NSObject
*/

type UILayoutGuide struct {
  *objc.NSObject

}

func (r *UILayoutGuide) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILayoutGuide) LeftAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILayoutGuide) WidthAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILayoutGuide) OwningView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILayoutGuide) LeadingAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILayoutGuide) SetIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILayoutGuide) TrailingAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILayoutGuide) RightAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILayoutGuide) TopAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILayoutGuide) HeightAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILayoutGuide) CenterXAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILayoutGuide) LayoutFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILayoutGuide) SetOwningView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILayoutGuide) CenterYAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILayoutGuide) BottomAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


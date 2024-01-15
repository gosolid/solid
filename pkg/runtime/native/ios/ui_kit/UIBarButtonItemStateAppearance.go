//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIBarButtonItemStateAppearance : objc.NSObject
*/

type UIBarButtonItemStateAppearance struct {
  *objc.NSObject

}

func (r *UIBarButtonItemStateAppearance) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemStateAppearance) TitleTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemStateAppearance) TitlePositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemStateAppearance) SetTitlePositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemStateAppearance) BackgroundImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemStateAppearance) SetBackgroundImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemStateAppearance) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemStateAppearance) BackgroundImagePositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemStateAppearance) SetBackgroundImagePositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemStateAppearance) SetTitleTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIWindow : UIKit.UIView
*/

type UIWindow struct {
  *UIView

}

func (r *UIWindow) WindowLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) CanBecomeKeyWindow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) RootViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) SetRootViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) SendEvent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) ConvertPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) Screen() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) SetWindowScene() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) CanResizeToFitContent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) SetCanResizeToFitContent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) InitWithWindowScene() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) SetScreen() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) WindowScene() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) SetWindowLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) IsKeyWindow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) BecomeKeyWindow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) MakeKeyWindow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) ConvertRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) ResignKeyWindow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindow) MakeKeyAndVisible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIMenuController : objc.NSObject
*/

type UIMenuController struct {
  *objc.NSObject

}

func (r *UIMenuController) MenuItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuController) SetMenuItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuController) SetMenuVisible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuController) HideMenuFromView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuController) HideMenu() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuController) ArrowDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuController) ShowMenuFromView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuController) Update() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuController) SharedMenuController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuController) SetArrowDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuController) SetTargetRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuController) IsMenuVisible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuController) MenuFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


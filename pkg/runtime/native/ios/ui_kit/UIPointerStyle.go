//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIPointerStyle : UIKit.UIHoverStyle
*/

type UIPointerStyle struct {
  *UIHoverStyle

}

func (r *UIPointerStyle) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerStyle) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerStyle) Accessories() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerStyle) SetAccessories() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerStyle) StyleWithEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerStyle) StyleWithShape() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerStyle) HiddenPointerStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerStyle) SystemPointerStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


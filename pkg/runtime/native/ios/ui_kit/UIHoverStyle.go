//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIHoverStyle : objc.NSObject
*/

type UIHoverStyle struct {
  *objc.NSObject

}

func (r *UIHoverStyle) SetShape() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverStyle) SetEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverStyle) StyleWithShape() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverStyle) AutomaticStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverStyle) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverStyle) Effect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverStyle) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverStyle) StyleWithEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverStyle) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverStyle) SetEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverStyle) Shape() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


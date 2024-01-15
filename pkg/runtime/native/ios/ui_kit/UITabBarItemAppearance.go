//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITabBarItemAppearance : objc.NSObject
*/

type UITabBarItemAppearance struct {
  *objc.NSObject

}

func (r *UITabBarItemAppearance) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemAppearance) InitWithStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemAppearance) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemAppearance) Copy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemAppearance) ConfigureWithDefaultForStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemAppearance) Normal() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemAppearance) Selected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemAppearance) Disabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemAppearance) Focused() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIEditMenuInteraction : objc.NSObject
*/

type UIEditMenuInteraction struct {
  *objc.NSObject

}

func (r *UIEditMenuInteraction) ReloadVisibleMenu() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEditMenuInteraction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEditMenuInteraction) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEditMenuInteraction) InitWithDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEditMenuInteraction) PresentEditMenuWithConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEditMenuInteraction) LocationInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEditMenuInteraction) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEditMenuInteraction) DismissMenu() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEditMenuInteraction) UpdateVisibleMenuPositionAnimated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


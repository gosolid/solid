//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIContextMenuInteraction : objc.NSObject
*/

type UIContextMenuInteraction struct {
  *objc.NSObject

}

func (r *UIContextMenuInteraction) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextMenuInteraction) MenuAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextMenuInteraction) InitWithDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextMenuInteraction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextMenuInteraction) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextMenuInteraction) LocationInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextMenuInteraction) UpdateVisibleMenuWithBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextMenuInteraction) DismissMenu() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIFocusSystem : objc.NSObject
*/

type UIFocusSystem struct {
  *objc.NSObject

}

func (r *UIFocusSystem) RequestFocusUpdateToEnvironment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusSystem) UpdateFocusIfNeeded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusSystem) Environment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusSystem) RegisterURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusSystem) FocusedItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusSystem) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusSystem) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusSystem) FocusSystemForEnvironment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


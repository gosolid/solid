//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPointerInteraction : objc.NSObject
*/

type UIPointerInteraction struct {
  *objc.NSObject

}

func (r *UIPointerInteraction) InitWithDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerInteraction) Invalidate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerInteraction) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerInteraction) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerInteraction) SetEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


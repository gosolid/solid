//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIDynamicBehavior : objc.NSObject
*/

type UIDynamicBehavior struct {
  *objc.NSObject

}

func (r *UIDynamicBehavior) RemoveChildBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicBehavior) WillMoveToAnimator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicBehavior) ChildBehaviors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicBehavior) Action() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicBehavior) SetAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicBehavior) DynamicAnimator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicBehavior) AddChildBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


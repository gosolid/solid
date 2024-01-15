//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIDynamicAnimator : objc.NSObject
*/

type UIDynamicAnimator struct {
  *objc.NSObject

}

func (r *UIDynamicAnimator) UpdateItemUsingCurrentState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicAnimator) ReferenceView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicAnimator) ElapsedTime() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicAnimator) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicAnimator) InitWithReferenceView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicAnimator) RemoveAllBehaviors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicAnimator) ItemsInRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicAnimator) Behaviors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicAnimator) IsRunning() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicAnimator) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicAnimator) AddBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicAnimator) RemoveBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


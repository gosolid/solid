//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIPushBehavior : UIKit.UIDynamicBehavior
*/

type UIPushBehavior struct {
  *UIDynamicBehavior

}

func (r *UIPushBehavior) InitWithItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPushBehavior) RemoveItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPushBehavior) Active() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPushBehavior) PushDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPushBehavior) Magnitude() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPushBehavior) SetTargetOffsetFromCenter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPushBehavior) SetAngle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPushBehavior) Mode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPushBehavior) SetActive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPushBehavior) TargetOffsetFromCenterForItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPushBehavior) SetPushDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPushBehavior) AddItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPushBehavior) Items() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPushBehavior) Angle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPushBehavior) SetMagnitude() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIGravityBehavior : UIKit.UIDynamicBehavior
*/

type UIGravityBehavior struct {
  *UIDynamicBehavior

}

func (r *UIGravityBehavior) SetGravityDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGravityBehavior) SetMagnitude() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGravityBehavior) InitWithItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGravityBehavior) AddItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGravityBehavior) RemoveItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGravityBehavior) GravityDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGravityBehavior) SetAngle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGravityBehavior) Items() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGravityBehavior) Angle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGravityBehavior) Magnitude() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


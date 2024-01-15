//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UISnapBehavior : UIKit.UIDynamicBehavior
*/

type UISnapBehavior struct {
  *UIDynamicBehavior

}

func (r *UISnapBehavior) SnapPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISnapBehavior) SetSnapPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISnapBehavior) Damping() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISnapBehavior) SetDamping() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISnapBehavior) InitWithItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIToolTipInteraction : objc.NSObject
*/

type UIToolTipInteraction struct {
  *objc.NSObject

}

func (r *UIToolTipInteraction) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIToolTipInteraction) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIToolTipInteraction) SetEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIToolTipInteraction) DefaultToolTip() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIToolTipInteraction) SetDefaultToolTip() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIToolTipInteraction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIToolTipInteraction) InitWithDefaultToolTip() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIToolTipInteraction) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


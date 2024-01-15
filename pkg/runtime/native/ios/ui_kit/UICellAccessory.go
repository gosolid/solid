//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UICellAccessory : objc.NSObject
*/

type UICellAccessory struct {
  *objc.NSObject

}

func (r *UICellAccessory) SetTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessory) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessory) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessory) DisplayedState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessory) SetReservedLayoutWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessory) TintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessory) SetDisplayedState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessory) IsHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessory) SetHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessory) ReservedLayoutWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


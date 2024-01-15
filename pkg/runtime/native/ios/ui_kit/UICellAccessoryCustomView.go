//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICellAccessoryCustomView : UIKit.UICellAccessory
*/

type UICellAccessoryCustomView struct {
  *UICellAccessory

}

func (r *UICellAccessoryCustomView) InitWithCustomView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryCustomView) CustomView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryCustomView) Placement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryCustomView) Position() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryCustomView) SetMaintainsFixedSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryCustomView) SetPosition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryCustomView) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryCustomView) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryCustomView) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryCustomView) MaintainsFixedSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


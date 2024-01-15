//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICellAccessoryLabel : UIKit.UICellAccessory
*/

type UICellAccessoryLabel struct {
  *UICellAccessory

}

func (r *UICellAccessoryLabel) InitWithText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryLabel) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryLabel) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryLabel) Text() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryLabel) Font() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryLabel) AdjustsFontForContentSizeCategory() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryLabel) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryLabel) SetFont() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryLabel) SetAdjustsFontForContentSizeCategory() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


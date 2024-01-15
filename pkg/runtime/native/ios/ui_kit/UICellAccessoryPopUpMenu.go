//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICellAccessoryPopUpMenu : UIKit.UICellAccessory
*/

type UICellAccessoryPopUpMenu struct {
  *UICellAccessory

}

func (r *UICellAccessoryPopUpMenu) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryPopUpMenu) Menu() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryPopUpMenu) SelectedElementDidChangeHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryPopUpMenu) SetSelectedElementDidChangeHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryPopUpMenu) InitWithMenu() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryPopUpMenu) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryPopUpMenu) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


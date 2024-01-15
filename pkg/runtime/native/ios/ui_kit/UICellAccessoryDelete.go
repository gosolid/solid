//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICellAccessoryDelete : UIKit.UICellAccessory
*/

type UICellAccessoryDelete struct {
  *UICellAccessory

}

func (r *UICellAccessoryDelete) ActionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryDelete) SetActionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryDelete) BackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryDelete) SetBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


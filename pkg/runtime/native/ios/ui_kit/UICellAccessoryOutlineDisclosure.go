//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICellAccessoryOutlineDisclosure : UIKit.UICellAccessory
*/

type UICellAccessoryOutlineDisclosure struct {
  *UICellAccessory

}

func (r *UICellAccessoryOutlineDisclosure) SetStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryOutlineDisclosure) ActionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryOutlineDisclosure) SetActionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryOutlineDisclosure) Style() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


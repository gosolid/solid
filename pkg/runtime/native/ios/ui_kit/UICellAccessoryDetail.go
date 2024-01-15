//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICellAccessoryDetail : UIKit.UICellAccessory
*/

type UICellAccessoryDetail struct {
  *UICellAccessory

}

func (r *UICellAccessoryDetail) ActionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryDetail) SetActionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICellAccessoryMultiselect : UIKit.UICellAccessory
*/

type UICellAccessoryMultiselect struct {
  *UICellAccessory

}

func (r *UICellAccessoryMultiselect) BackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryMultiselect) SetBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


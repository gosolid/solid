//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICellAccessoryReorder : UIKit.UICellAccessory
*/

type UICellAccessoryReorder struct {
  *UICellAccessory

}

func (r *UICellAccessoryReorder) ShowsVerticalSeparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryReorder) SetShowsVerticalSeparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICellAccessoryInsert : UIKit.UICellAccessory
*/

type UICellAccessoryInsert struct {
  *UICellAccessory

}

func (r *UICellAccessoryInsert) BackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryInsert) SetBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryInsert) ActionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICellAccessoryInsert) SetActionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


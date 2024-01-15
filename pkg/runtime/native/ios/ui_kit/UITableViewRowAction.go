//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITableViewRowAction : objc.NSObject
*/

type UITableViewRowAction struct {
  *objc.NSObject

}

func (r *UITableViewRowAction) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewRowAction) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewRowAction) BackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewRowAction) SetBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewRowAction) BackgroundEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewRowAction) SetBackgroundEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewRowAction) RowActionWithStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewRowAction) Style() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITableViewPlaceholder : objc.NSObject
*/

type UITableViewPlaceholder struct {
  *objc.NSObject

}

func (r *UITableViewPlaceholder) InitWithInsertionIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewPlaceholder) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewPlaceholder) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewPlaceholder) CellUpdateHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewPlaceholder) SetCellUpdateHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


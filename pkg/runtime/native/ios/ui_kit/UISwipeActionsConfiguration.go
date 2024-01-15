//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISwipeActionsConfiguration : objc.NSObject
*/

type UISwipeActionsConfiguration struct {
  *objc.NSObject

}

func (r *UISwipeActionsConfiguration) PerformsFirstActionWithFullSwipe() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISwipeActionsConfiguration) SetPerformsFirstActionWithFullSwipe() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISwipeActionsConfiguration) ConfigurationWithActions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISwipeActionsConfiguration) Actions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITableViewFocusUpdateContext : UIKit.UIFocusUpdateContext
*/

type UITableViewFocusUpdateContext struct {
  *UIFocusUpdateContext

}

func (r *UITableViewFocusUpdateContext) PreviouslyFocusedIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewFocusUpdateContext) NextFocusedIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


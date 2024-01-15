//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITableViewDropPlaceholder : UIKit.UITableViewPlaceholder
*/

type UITableViewDropPlaceholder struct {
  *UITableViewPlaceholder

}

func (r *UITableViewDropPlaceholder) PreviewParametersProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewDropPlaceholder) SetPreviewParametersProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


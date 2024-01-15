//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITargetedDragPreview : UIKit.UITargetedPreview
*/

type UITargetedDragPreview struct {
  *UITargetedPreview

}

func (r *UITargetedDragPreview) RetargetedPreviewWithTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


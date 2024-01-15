//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITargetedPreview : objc.NSObject
*/

type UITargetedPreview struct {
  *objc.NSObject

}

func (r *UITargetedPreview) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITargetedPreview) RetargetedPreviewWithTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITargetedPreview) Target() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITargetedPreview) View() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITargetedPreview) Parameters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITargetedPreview) Size() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITargetedPreview) InitWithView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITargetedPreview) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


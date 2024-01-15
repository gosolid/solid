//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UILargeContentViewerInteraction : objc.NSObject
*/

type UILargeContentViewerInteraction struct {
  *objc.NSObject

}

func (r *UILargeContentViewerInteraction) InitWithDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILargeContentViewerInteraction) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILargeContentViewerInteraction) GestureRecognizerForExclusionRelationship() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILargeContentViewerInteraction) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


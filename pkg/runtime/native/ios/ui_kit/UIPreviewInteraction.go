//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPreviewInteraction : objc.NSObject
*/

type UIPreviewInteraction struct {
  *objc.NSObject

}

func (r *UIPreviewInteraction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewInteraction) LocationInCoordinateSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewInteraction) CancelInteraction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewInteraction) View() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewInteraction) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewInteraction) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewInteraction) InitWithView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


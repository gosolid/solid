//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIFocusUpdateContext : objc.NSObject
*/

type UIFocusUpdateContext struct {
  *objc.NSObject

}

func (r *UIFocusUpdateContext) PreviouslyFocusedItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusUpdateContext) NextFocusedItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusUpdateContext) PreviouslyFocusedView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusUpdateContext) NextFocusedView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusUpdateContext) FocusHeading() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


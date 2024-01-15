//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIContextualAction : objc.NSObject
*/

type UIContextualAction struct {
  *objc.NSObject

}

func (r *UIContextualAction) ContextualActionWithStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextualAction) Handler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextualAction) Image() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextualAction) SetImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextualAction) Style() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextualAction) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextualAction) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextualAction) BackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextualAction) SetBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


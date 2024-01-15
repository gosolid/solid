//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITextSelectionRect : objc.NSObject
*/

type UITextSelectionRect struct {
  *objc.NSObject

}

func (r *UITextSelectionRect) IsVertical() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionRect) Rect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionRect) WritingDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionRect) ContainsStart() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextSelectionRect) ContainsEnd() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


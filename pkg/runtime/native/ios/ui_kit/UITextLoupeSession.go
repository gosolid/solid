//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITextLoupeSession : objc.NSObject
*/

type UITextLoupeSession struct {
  *objc.NSObject

}

func (r *UITextLoupeSession) BeginLoupeSessionAtPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextLoupeSession) MoveToPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextLoupeSession) Invalidate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


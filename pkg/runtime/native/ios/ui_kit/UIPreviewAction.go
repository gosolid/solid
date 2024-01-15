//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPreviewAction : objc.NSObject
*/

type UIPreviewAction struct {
  *objc.NSObject

}

func (r *UIPreviewAction) ActionWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewAction) Handler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


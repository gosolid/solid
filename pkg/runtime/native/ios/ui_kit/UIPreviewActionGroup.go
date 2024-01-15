//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPreviewActionGroup : objc.NSObject
*/

type UIPreviewActionGroup struct {
  *objc.NSObject

}

func (r *UIPreviewActionGroup) ActionGroupWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


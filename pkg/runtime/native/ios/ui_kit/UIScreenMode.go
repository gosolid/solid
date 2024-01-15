//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIScreenMode : objc.NSObject
*/

type UIScreenMode struct {
  *objc.NSObject

}

func (r *UIScreenMode) PixelAspectRatio() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreenMode) Size() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


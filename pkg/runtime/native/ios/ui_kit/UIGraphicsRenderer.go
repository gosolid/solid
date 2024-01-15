//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIGraphicsRenderer : objc.NSObject
*/

type UIGraphicsRenderer struct {
  *objc.NSObject

}

func (r *UIGraphicsRenderer) InitWithBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsRenderer) Format() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsRenderer) AllowsImageOutput() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIGraphicsRendererFormat : objc.NSObject
*/

type UIGraphicsRendererFormat struct {
  *objc.NSObject

}

func (r *UIGraphicsRendererFormat) DefaultFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsRendererFormat) PreferredFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsRendererFormat) Bounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


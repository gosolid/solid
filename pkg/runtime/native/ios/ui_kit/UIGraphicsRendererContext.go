//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIGraphicsRendererContext : objc.NSObject
*/

type UIGraphicsRendererContext struct {
  *objc.NSObject

}

func (r *UIGraphicsRendererContext) StrokeRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsRendererContext) ClipToRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsRendererContext) CGContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsRendererContext) Format() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGraphicsRendererContext) FillRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


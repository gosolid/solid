//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIShapeResolutionContext : objc.NSObject
*/

type UIShapeResolutionContext struct {
  *objc.NSObject

}

func (r *UIShapeResolutionContext) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIShapeResolutionContext) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIShapeResolutionContext) ContentShape() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPointerShape : objc.NSObject
*/

type UIPointerShape struct {
  *objc.NSObject

}

func (r *UIPointerShape) ShapeWithPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerShape) ShapeWithRoundedRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerShape) BeamWithPreferredLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerShape) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerShape) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


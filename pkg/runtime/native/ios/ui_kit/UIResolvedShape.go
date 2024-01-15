//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIResolvedShape : objc.NSObject
*/

type UIResolvedShape struct {
  *objc.NSObject

}

func (r *UIResolvedShape) ShapeByApplyingInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResolvedShape) ShapeByApplyingInset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResolvedShape) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResolvedShape) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResolvedShape) Shape() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResolvedShape) BoundingRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResolvedShape) Path() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


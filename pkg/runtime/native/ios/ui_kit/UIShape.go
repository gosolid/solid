//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIShape : objc.NSObject
*/

type UIShape struct {
  *objc.NSObject

}

func (r *UIShape) ShapeByApplyingInset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIShape) ResolvedShapeInContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIShape) RectShape() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIShape) CapsuleShape() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIShape) RectShapeWithCornerRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIShape) ShapeWithProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIShape) ShapeByApplyingInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIShape) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIShape) CircleShape() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIShape) FixedRectShapeWithRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIShape) ShapeWithBezierPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIShape) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


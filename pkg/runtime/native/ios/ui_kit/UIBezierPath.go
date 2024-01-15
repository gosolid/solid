//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIBezierPath : objc.NSObject
*/

type UIBezierPath struct {
  *objc.NSObject

}

func (r *UIBezierPath) ContainsPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) GetLineDash() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) SetLineJoinStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) LineJoinStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) MiterLimit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) FillWithBlendMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) AddClip() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) IsEmpty() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) BezierPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) AddLineToPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) RemoveAllPoints() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) SetLineWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) LineCapStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) UsesEvenOddFillRule() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) BezierPathWithOvalInRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) MoveToPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) Bounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) LineWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) Flatness() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) BezierPathWithRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) Stroke() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) SetCGPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) CurrentPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) BezierPathWithCGPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) AddQuadCurveToPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) ClosePath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) Fill() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) CGPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) AddCurveToPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) AddArcWithCenter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) StrokeWithBlendMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) SetFlatness() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) SetLineDash() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) SetLineCapStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) SetMiterLimit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) BezierPathWithRoundedRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) BezierPathWithArcCenter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) AppendPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) BezierPathByReversingPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) ApplyTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBezierPath) SetUsesEvenOddFillRule() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
)

/*
interface QuartzCore.CAShapeLayer : QuartzCore.CALayer
*/

type CAShapeLayer struct {
  *CALayer

}

func (r *CAShapeLayer) SetLineJoin() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) SetLineDashPhase() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) StrokeEnd() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) SetMiterLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) LineCap() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) SetLineCap() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) FillRule() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) SetStrokeEnd() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) SetFillRule() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) StrokeColor() (*core_graphics.CGColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) SetStrokeColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) StrokeStart() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) SetStrokeStart() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) SetLineWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) SetPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) SetFillColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) LineDashPattern() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) SetLineDashPattern() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) LineWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) MiterLimit() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) LineJoin() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) LineDashPhase() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) Path() (*core_graphics.CGPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAShapeLayer) FillColor() (*core_graphics.CGColor, error) {
  return nil, fmt.Errorf("unimplemented")
}


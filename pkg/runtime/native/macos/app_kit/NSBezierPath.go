//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
)

/*
interface AppKit.NSBezierPath : objc.NSObject
*/

type NSBezierPath struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *NSBezierPath) SetWindingRule() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) ElementCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) AppendBezierPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) AppendBezierPathWithArcFromPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) DefaultLineJoinStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) RelativeMoveToPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) RelativeCurveToPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetDefaultLineWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) ClipRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) DefaultLineWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) BezierPath() (*NSBezierPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) BezierPathWithOvalInRect() (*NSBezierPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) StrokeRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) LineJoinStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) ClosePath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) DefaultMiterLimit() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) DefaultWindingRule() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) AppendBezierPathWithCGGlyphs() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetDefaultLineJoinStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) MoveToPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) RemoveAllPoints() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) AppendBezierPathWithArcWithCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) LineWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) CurrentPoint() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) ControlPointBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetClip() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) AppendBezierPathWithCGGlyph() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetDefaultFlatness() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetAssociatedPoints() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) IsEmpty() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) FillRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) StrokeLineFromPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) GetLineDash() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetLineJoinStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) LineToPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetDefaultMiterLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetDefaultLineCapStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) TransformUsingAffineTransform() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) AppendBezierPathWithPoints() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) Stroke() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) Fill() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) AppendBezierPathWithRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetDefaultWindingRule() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetFlatness() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) BezierPathWithRect() (*NSBezierPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) BezierPathWithRoundedRect() (*NSBezierPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) RelativeLineToPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) BezierPathByFlatteningPath() (*NSBezierPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) BezierPathByReversingPath() (*NSBezierPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetLineWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) LineCapStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) MiterLimit() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) Flatness() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) CurveToPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) ElementAtIndex() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetCGPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetLineCapStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) WindingRule() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) AddClip() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) CGPath() (*core_graphics.CGPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) DefaultFlatness() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) AppendBezierPathWithRoundedRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetMiterLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) BezierPathWithCGPath() (*NSBezierPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) SetLineDash() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) AppendBezierPathWithOvalInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) DefaultLineCapStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) Bounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) DrawPackedGlyphs() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBezierPath) ContainsPoint() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


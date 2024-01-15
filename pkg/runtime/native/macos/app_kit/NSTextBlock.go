//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSTextBlock : objc.NSObject
*/

type NSTextBlock struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *NSTextBlock) ValueForDimension() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) WidthForLayer() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) WidthValueTypeForLayer() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) SetValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) SetBorderColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) RectForLayoutAtPoint() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) BoundsRectForContentRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) ContentWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) VerticalAlignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) ValueTypeForDimension() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) SetContentWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) ContentWidthValueType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) BorderColorForEdge() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) DrawBackgroundWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextBlock) SetVerticalAlignment() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSTextLineFragment : objc.NSObject
*/

type NSTextLineFragment struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *NSTextLineFragment) AttributedString() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) LocationForCharacterAtIndex() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) CharacterIndexForPoint() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) FractionOfDistanceThroughGlyphForPoint() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) CharacterRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) TypographicBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) GlyphOrigin() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) InitWithAttributedString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) InitWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) DrawAtPoint() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSTextLineFragment : objc.NSObject
*/

type NSTextLineFragment struct {
  *objc.NSObject

}

func (r *NSTextLineFragment) InitWithString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) DrawAtPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) LocationForCharacterAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) FractionOfDistanceThroughGlyphForPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) AttributedString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) GlyphOrigin() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) InitWithAttributedString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) CharacterIndexForPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) CharacterRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextLineFragment) TypographicBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


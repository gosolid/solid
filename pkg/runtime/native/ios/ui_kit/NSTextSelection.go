//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSTextSelection : objc.NSObject
*/

type NSTextSelection struct {
  *objc.NSObject

}

func (r *NSTextSelection) SetTypingAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) IsLogical() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) InitWithRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) InitWithLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) TextSelectionWithTextRanges() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) Affinity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) AnchorPositionOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) SetLogical() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) InitWithRanges() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) TextRanges() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) SetAnchorPositionOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) SecondarySelectionLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) SetSecondarySelectionLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) TypingAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) IsTransient() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) Granularity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


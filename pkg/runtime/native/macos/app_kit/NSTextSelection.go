//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSTextSelection : objc.NSObject
*/

type NSTextSelection struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *NSTextSelection) SetTypingAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) InitWithLocation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) SecondarySelectionLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) TextSelectionWithTextRanges() (*NSTextSelection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) Affinity() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) AnchorPositionOffset() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) SetAnchorPositionOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) IsLogical() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) TypingAttributes() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) InitWithRanges() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) InitWithRange() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) SetLogical() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) Granularity() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) IsTransient() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) SetSecondarySelectionLocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelection) TextRanges() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


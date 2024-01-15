//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSTextInputClient
*/

type NSTextInputClient struct {

}

func (r *NSTextInputClient) PreferredTextAccessoryPlacement() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) HasMarkedText() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) FractionOfDistanceThroughGlyphForPoint() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) DrawsVerticallyForCharacterAtIndex() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) MarkedRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) WindowLevel() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) SelectedRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) ValidAttributesForMarkedText() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) CharacterIndexForPoint() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) AttributedString() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) BaselineDeltaForCharacterAtIndex() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) InsertText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) SetMarkedText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) UnmarkText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) UnionRectInVisibleSelectedRange() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) DocumentVisibleRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) DoCommandBySelector() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) AttributedSubstringForProposedRange() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextInputClient) FirstRectForCharacterRange() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}


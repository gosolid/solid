//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
protocol AppKit.NSTextInput
*/

type NSTextInput struct {

}

func (r *NSTextInput) FirstRectForCharacterRange() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextInput) CharacterIndexForPoint() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInput) DoCommandBySelector() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInput) MarkedRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSTextInput) SelectedRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSTextInput) HasMarkedText() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextInput) ConversationIdentifier() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInput) AttributedSubstringFromRange() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextInput) ValidAttributesForMarkedText() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextInput) InsertText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInput) SetMarkedText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInput) UnmarkText() error {
  return fmt.Errorf("unimplemented")
}


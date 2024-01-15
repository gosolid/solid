//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSTextFinderClient
*/

type NSTextFinderClient struct {

}

func (r *NSTextFinderClient) ScrollRangeToVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) ShouldReplaceCharactersInRanges() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) DidReplaceCharacters() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) AllowsMultipleSelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) SetSelectedRanges() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) StringLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) RectsForCharacterRange() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) DrawCharactersInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) IsEditable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) FirstSelectedRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) StringAtIndex() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) ReplaceCharactersInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) ContentViewAtIndex() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) String() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) VisibleCharacterRanges() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) IsSelectable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextFinderClient) SelectedRanges() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


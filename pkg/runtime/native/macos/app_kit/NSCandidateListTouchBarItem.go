//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSCandidateListTouchBarItem : AppKit.NSTouchBarItem
*/

type NSCandidateListTouchBarItem struct {
  *NSTouchBarItem

}

func (r *NSCandidateListTouchBarItem) SetCollapsed() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) IsCandidateListVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) Candidates() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) CustomizationLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) Client() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) AllowsCollapsing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) SetAllowsCollapsing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) SetAllowsTextInputContextCandidates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) SetAttributedStringForCandidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) IsCollapsed() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) AllowsTextInputContextCandidates() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) AttributedStringForCandidate() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) SetCustomizationLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) UpdateWithInsertionPointVisibility() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) SetCandidates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCandidateListTouchBarItem) SetClient() error {
  return fmt.Errorf("unimplemented")
}


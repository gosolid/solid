//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSTextFinder : objc.NSObject
*/

type NSTextFinder struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSTextFinder) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) FindBarContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) SetFindBarContainer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) SetFindIndicatorNeedsUpdate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) SetIncrementalSearchingShouldDimContentView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) CancelFindIndicator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) DrawIncrementalMatchHighlightInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) NoteClientStringWillChange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) Client() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) IncrementalSearchingShouldDimContentView() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) IncrementalMatchRanges() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) PerformAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) SetClient() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) FindIndicatorNeedsUpdate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) IsIncrementalSearchingEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) SetIncrementalSearchingEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFinder) ValidateAction() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


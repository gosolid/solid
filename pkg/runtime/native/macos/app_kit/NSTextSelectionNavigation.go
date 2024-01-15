//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTextSelectionNavigation : objc.NSObject
*/

type NSTextSelectionNavigation struct {
  *objc.NSObject

}

func (r *NSTextSelectionNavigation) DestinationSelectionForTextSelection() (*NSTextSelection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) SetAllowsNonContiguousRanges() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) RotatesCoordinateSystemForLayoutOrientation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) DeletionRangesForTextSelection() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) TextSelectionsInteractingAtPoint() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) TextSelectionForSelectionGranularity() (*NSTextSelection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) AllowsNonContiguousRanges() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) SetRotatesCoordinateSystemForLayoutOrientation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) FlushLayoutCache() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) TextSelectionDataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) InitWithDataSource() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) ResolvedInsertionLocationForTextSelection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


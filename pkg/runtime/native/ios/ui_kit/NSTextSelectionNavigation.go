//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSTextSelectionNavigation : objc.NSObject
*/

type NSTextSelectionNavigation struct {
  *objc.NSObject

}

func (r *NSTextSelectionNavigation) InitWithDataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) TextSelectionForSelectionGranularity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) RotatesCoordinateSystemForLayoutOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) FlushLayoutCache() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) TextSelectionsInteractingAtPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) SetAllowsNonContiguousRanges() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) AllowsNonContiguousRanges() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) DestinationSelectionForTextSelection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) DeletionRangesForTextSelection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) TextSelectionDataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) ResolvedInsertionLocationForTextSelection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextSelectionNavigation) SetRotatesCoordinateSystemForLayoutOrientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSGridColumn : objc.NSObject
*/

type NSGridColumn struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSGridColumn) NumberOfCells() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridColumn) XPlacement() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridColumn) Width() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridColumn) SetLeadingPadding() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridColumn) TrailingPadding() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridColumn) MergeCellsInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridColumn) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridColumn) GridView() (*NSGridView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridColumn) SetTrailingPadding() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridColumn) IsHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGridColumn) SetHidden() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridColumn) CellAtIndex() (*NSGridCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridColumn) SetXPlacement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridColumn) LeadingPadding() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}


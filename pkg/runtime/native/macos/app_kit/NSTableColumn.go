//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSTableColumn : objc.NSObject
*/

type NSTableColumn struct {
  *objc.NSObject
  *foundation.NSCoding
  *NSUserInterfaceItemIdentification
}

func (r *NSTableColumn) SetMinWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) IsHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) SetTableView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) MinWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) HeaderCell() (*NSTableHeaderCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) SortDescriptorPrototype() (*foundation.NSSortDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) SetHidden() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) InitWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) ResizingMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) SetResizingMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) SizeToFit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) TableView() (*NSTableView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) Width() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) MaxWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) SetHeaderCell() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) HeaderToolTip() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) Identifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) SetIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) SetEditable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) SetSortDescriptorPrototype() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) SetHeaderToolTip() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) SetMaxWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableColumn) IsEditable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSGridCell : objc.NSObject
*/

type NSGridCell struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSGridCell) SetRowAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridCell) SetCustomPlacementConstraints() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridCell) SetContentView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridCell) Row() (*NSGridRow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridCell) Column() (*NSGridColumn, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridCell) SetXPlacement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridCell) SetYPlacement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridCell) RowAlignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridCell) CustomPlacementConstraints() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridCell) ContentView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridCell) EmptyContentView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridCell) XPlacement() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridCell) YPlacement() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


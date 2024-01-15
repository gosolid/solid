//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSTableRowView : AppKit.NSView
*/

type NSTableRowView struct {
  *NSView
  *NSAccessibilityRow
}

func (r *NSTableRowView) IsGroupRowStyle() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) SetGroupRowStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) IsSelected() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) SetSelected() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) IsNextRowSelected() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) ViewAtColumn() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) DrawSeparatorInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) DrawDraggingDestinationFeedbackInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) SetPreviousRowSelected() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) SetTargetForDropOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) SetIndentationForDropOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) NumberOfColumns() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) DrawSelectionInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) IsFloating() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) IndentationForDropOperation() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) IsPreviousRowSelected() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) IsEmphasized() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) IsTargetForDropOperation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) InteriorBackgroundStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) DrawBackgroundInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) SetSelectionHighlightStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) SetNextRowSelected() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) SetFloating() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) DraggingDestinationFeedbackStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) SetDraggingDestinationFeedbackStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) SelectionHighlightStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableRowView) SetEmphasized() error {
  return fmt.Errorf("unimplemented")
}


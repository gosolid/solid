//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTableCellView : AppKit.NSView
*/

type NSTableCellView struct {
  *NSView

}

func (r *NSTableCellView) ObjectValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableCellView) SetObjectValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableCellView) ImageView() (*NSImageView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableCellView) RowSizeStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableCellView) TextField() (*NSTextField, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableCellView) SetTextField() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableCellView) SetImageView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableCellView) BackgroundStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableCellView) SetBackgroundStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableCellView) SetRowSizeStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableCellView) DraggingImageComponents() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


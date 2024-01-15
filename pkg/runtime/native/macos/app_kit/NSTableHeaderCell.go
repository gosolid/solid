//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSTableHeaderCell : AppKit.NSTextFieldCell
*/

type NSTableHeaderCell struct {
  *NSTextFieldCell

}

func (r *NSTableHeaderCell) DrawSortIndicatorWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableHeaderCell) SortIndicatorRectForBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}


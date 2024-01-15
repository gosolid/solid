//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSTextTable : AppKit.NSTextBlock
*/

type NSTextTable struct {
  *NSTextBlock

}

func (r *NSTextTable) BoundsRectForBlock() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextTable) LayoutAlgorithm() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextTable) CollapsesBorders() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextTable) SetCollapsesBorders() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextTable) HidesEmptyCells() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextTable) SetHidesEmptyCells() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextTable) RectForBlock() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTextTable) DrawBackgroundForBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextTable) NumberOfColumns() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextTable) SetNumberOfColumns() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextTable) SetLayoutAlgorithm() error {
  return fmt.Errorf("unimplemented")
}


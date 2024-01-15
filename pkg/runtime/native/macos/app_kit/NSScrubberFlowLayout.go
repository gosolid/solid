//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSScrubberFlowLayout : AppKit.NSScrubberLayout
*/

type NSScrubberFlowLayout struct {
  *NSScrubberLayout

}

func (r *NSScrubberFlowLayout) ItemSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSScrubberFlowLayout) SetItemSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberFlowLayout) InvalidateLayoutForItemsAtIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberFlowLayout) ItemSpacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrubberFlowLayout) SetItemSpacing() error {
  return fmt.Errorf("unimplemented")
}


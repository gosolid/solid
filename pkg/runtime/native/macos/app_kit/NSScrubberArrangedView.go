//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSScrubberArrangedView : AppKit.NSView
*/

type NSScrubberArrangedView struct {
  *NSView

}

func (r *NSScrubberArrangedView) ApplyLayoutAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberArrangedView) IsSelected() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrubberArrangedView) SetSelected() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberArrangedView) IsHighlighted() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrubberArrangedView) SetHighlighted() error {
  return fmt.Errorf("unimplemented")
}


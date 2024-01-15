//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSTextInsertionIndicator : AppKit.NSView
*/

type NSTextInsertionIndicator struct {
  *NSView

}

func (r *NSTextInsertionIndicator) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInsertionIndicator) AutomaticModeOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInsertionIndicator) SetAutomaticModeOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInsertionIndicator) EffectsViewInserter() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextInsertionIndicator) SetEffectsViewInserter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInsertionIndicator) DisplayMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextInsertionIndicator) SetDisplayMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextInsertionIndicator) Color() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSColorPickingCustom
*/

type NSColorPickingCustom struct {

}

func (r *NSColorPickingCustom) SupportsMode() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSColorPickingCustom) CurrentMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColorPickingCustom) ProvideNewView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPickingCustom) SetColor() error {
  return fmt.Errorf("unimplemented")
}


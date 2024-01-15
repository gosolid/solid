//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSMenuDelegate
*/

type NSMenuDelegate struct {

}

func (r *NSMenuDelegate) MenuDidClose() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuDelegate) ConfinementRectForMenu() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSMenuDelegate) MenuNeedsUpdate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuDelegate) NumberOfItemsInMenu() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenuDelegate) Menu() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuDelegate) MenuHasKeyEquivalent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuDelegate) MenuWillOpen() error {
  return fmt.Errorf("unimplemented")
}


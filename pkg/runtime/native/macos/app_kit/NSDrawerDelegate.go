//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSDrawerDelegate
*/

type NSDrawerDelegate struct {

}

func (r *NSDrawerDelegate) DrawerShouldClose() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDrawerDelegate) DrawerWillResizeContents() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSDrawerDelegate) DrawerWillOpen() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawerDelegate) DrawerDidOpen() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawerDelegate) DrawerWillClose() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawerDelegate) DrawerDidClose() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDrawerDelegate) DrawerShouldOpen() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


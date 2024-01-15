//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSMediaLibraryBrowserController : objc.NSObject
*/

type NSMediaLibraryBrowserController struct {
  *objc.NSObject

}

func (r *NSMediaLibraryBrowserController) Frame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSMediaLibraryBrowserController) SetFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMediaLibraryBrowserController) MediaLibraries() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMediaLibraryBrowserController) SetMediaLibraries() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMediaLibraryBrowserController) TogglePanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMediaLibraryBrowserController) SharedMediaLibraryBrowserController() (*NSMediaLibraryBrowserController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMediaLibraryBrowserController) IsVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMediaLibraryBrowserController) SetVisible() error {
  return fmt.Errorf("unimplemented")
}


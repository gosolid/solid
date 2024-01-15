//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSStatusItem : objc.NSObject
*/

type NSStatusItem struct {
  *objc.NSObject

}

func (r *NSStatusItem) Menu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStatusItem) SetMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStatusItem) Behavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStatusItem) SetVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStatusItem) AutosaveName() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStatusItem) StatusBar() (*NSStatusBar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStatusItem) Length() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStatusItem) SetLength() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStatusItem) Button() (*NSStatusBarButton, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStatusItem) SetBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStatusItem) IsVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSStatusItem) SetAutosaveName() error {
  return fmt.Errorf("unimplemented")
}


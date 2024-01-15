//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSWindowTabGroup : objc.NSObject
*/

type NSWindowTabGroup struct {
  *objc.NSObject

}

func (r *NSWindowTabGroup) InsertWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowTabGroup) RemoveWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowTabGroup) IsOverviewVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindowTabGroup) SetOverviewVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowTabGroup) SelectedWindow() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowTabGroup) AddWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowTabGroup) Identifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowTabGroup) Windows() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowTabGroup) IsTabBarVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindowTabGroup) SetSelectedWindow() error {
  return fmt.Errorf("unimplemented")
}


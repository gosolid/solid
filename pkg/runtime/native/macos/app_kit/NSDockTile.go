//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSDockTile : objc.NSObject
*/

type NSDockTile struct {
  *objc.NSObject

}

func (r *NSDockTile) Size() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSDockTile) ContentView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDockTile) SetContentView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDockTile) BadgeLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDockTile) Display() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDockTile) SetShowsApplicationBadge() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDockTile) SetBadgeLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDockTile) Owner() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDockTile) ShowsApplicationBadge() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


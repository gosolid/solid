//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSDockTilePlugIn
*/

type NSDockTilePlugIn struct {

}

func (r *NSDockTilePlugIn) SetDockTile() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDockTilePlugIn) DockMenu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}


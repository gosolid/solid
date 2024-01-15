//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSTrackingSeparatorToolbarItem : AppKit.NSToolbarItem
*/

type NSTrackingSeparatorToolbarItem struct {
  *NSToolbarItem

}

func (r *NSTrackingSeparatorToolbarItem) TrackingSeparatorToolbarItemWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTrackingSeparatorToolbarItem) SplitView() (*NSSplitView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTrackingSeparatorToolbarItem) SetSplitView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTrackingSeparatorToolbarItem) DividerIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTrackingSeparatorToolbarItem) SetDividerIndex() error {
  return fmt.Errorf("unimplemented")
}


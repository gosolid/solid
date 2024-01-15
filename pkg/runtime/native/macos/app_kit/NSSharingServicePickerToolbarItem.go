//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSSharingServicePickerToolbarItem : AppKit.NSToolbarItem
*/

type NSSharingServicePickerToolbarItem struct {
  *NSToolbarItem

}

func (r *NSSharingServicePickerToolbarItem) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingServicePickerToolbarItem) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}


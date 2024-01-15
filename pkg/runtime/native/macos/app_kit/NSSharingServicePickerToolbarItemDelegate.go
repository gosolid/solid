//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSSharingServicePickerToolbarItemDelegate
*/

type NSSharingServicePickerToolbarItemDelegate struct {

}

func (r *NSSharingServicePickerToolbarItemDelegate) ItemsForSharingServicePickerToolbarItem() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


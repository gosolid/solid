//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSSharingServicePickerTouchBarItemDelegate
*/

type NSSharingServicePickerTouchBarItemDelegate struct {

}

func (r *NSSharingServicePickerTouchBarItemDelegate) ItemsForSharingServicePickerTouchBarItem() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSSharingServicePickerDelegate
*/

type NSSharingServicePickerDelegate struct {

}

func (r *NSSharingServicePickerDelegate) SharingServicePicker() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSSharingServicePickerTouchBarItem : AppKit.NSTouchBarItem
*/

type NSSharingServicePickerTouchBarItem struct {
  *NSTouchBarItem

}

func (r *NSSharingServicePickerTouchBarItem) ButtonImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingServicePickerTouchBarItem) SetButtonImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSharingServicePickerTouchBarItem) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingServicePickerTouchBarItem) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSharingServicePickerTouchBarItem) IsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSharingServicePickerTouchBarItem) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSharingServicePickerTouchBarItem) ButtonTitle() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingServicePickerTouchBarItem) SetButtonTitle() error {
  return fmt.Errorf("unimplemented")
}


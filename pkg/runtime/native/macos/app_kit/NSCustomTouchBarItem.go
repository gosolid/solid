//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSCustomTouchBarItem : AppKit.NSTouchBarItem
*/

type NSCustomTouchBarItem struct {
  *NSTouchBarItem

}

func (r *NSCustomTouchBarItem) CustomizationLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCustomTouchBarItem) SetCustomizationLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCustomTouchBarItem) View() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCustomTouchBarItem) SetView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCustomTouchBarItem) ViewController() (*NSViewController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCustomTouchBarItem) SetViewController() error {
  return fmt.Errorf("unimplemented")
}


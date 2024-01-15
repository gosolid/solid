//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSGroupTouchBarItem : AppKit.NSTouchBarItem
*/

type NSGroupTouchBarItem struct {
  *NSTouchBarItem

}

func (r *NSGroupTouchBarItem) SetPrioritizedCompressionOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGroupTouchBarItem) GroupItemWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGroupTouchBarItem) AlertStyleGroupItemWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGroupTouchBarItem) SetGroupTouchBar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGroupTouchBarItem) SetCustomizationLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGroupTouchBarItem) GroupUserInterfaceLayoutDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGroupTouchBarItem) PrioritizedCompressionOptions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGroupTouchBarItem) GroupTouchBar() (*NSTouchBar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGroupTouchBarItem) SetGroupUserInterfaceLayoutDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGroupTouchBarItem) SetPrefersEqualWidths() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGroupTouchBarItem) CustomizationLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGroupTouchBarItem) SetPreferredItemWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGroupTouchBarItem) EffectiveCompressionOptions() (*NSUserInterfaceCompressionOptions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGroupTouchBarItem) PrefersEqualWidths() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGroupTouchBarItem) PreferredItemWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}


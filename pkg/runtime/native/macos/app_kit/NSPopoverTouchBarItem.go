//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSPopoverTouchBarItem : AppKit.NSTouchBarItem
*/

type NSPopoverTouchBarItem struct {
  *NSTouchBarItem

}

func (r *NSPopoverTouchBarItem) SetCollapsedRepresentationImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) PressAndHoldTouchBar() (*NSTouchBar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) SetShowsCloseButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) MakeStandardActivatePopoverGestureRecognizer() (*NSGestureRecognizer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) SetPopoverTouchBar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) CollapsedRepresentation() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) SetCollapsedRepresentation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) PopoverTouchBar() (*NSTouchBar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) ShowsCloseButton() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) CustomizationLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) SetCustomizationLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) SetCollapsedRepresentationLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) SetPressAndHoldTouchBar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) ShowPopover() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) DismissPopover() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) CollapsedRepresentationImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopoverTouchBarItem) CollapsedRepresentationLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}


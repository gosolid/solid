//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSPopoverDelegate
*/

type NSPopoverDelegate struct {

}

func (r *NSPopoverDelegate) PopoverDidDetach() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopoverDelegate) DetachableWindowForPopover() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopoverDelegate) PopoverWillShow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopoverDelegate) PopoverDidShow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopoverDelegate) PopoverWillClose() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopoverDelegate) PopoverDidClose() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopoverDelegate) PopoverShouldClose() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPopoverDelegate) PopoverShouldDetach() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTableViewRowAction : objc.NSObject
*/

type NSTableViewRowAction struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSTableViewRowAction) RowActionWithStyle() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableViewRowAction) Style() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableViewRowAction) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableViewRowAction) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableViewRowAction) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableViewRowAction) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableViewRowAction) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableViewRowAction) SetImage() error {
  return fmt.Errorf("unimplemented")
}


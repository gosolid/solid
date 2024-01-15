//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSWindowTab : objc.NSObject
*/

type NSWindowTab struct {
  *objc.NSObject

}

func (r *NSWindowTab) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowTab) AttributedTitle() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowTab) SetAttributedTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowTab) ToolTip() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowTab) SetToolTip() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowTab) AccessoryView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowTab) SetAccessoryView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowTab) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}


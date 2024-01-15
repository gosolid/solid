//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSTabViewItem : objc.NSObject
*/

type NSTabViewItem struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSTabViewItem) InitialFirstResponder() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) InitWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) View() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) SetView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) SetToolTip() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) SizeOfLabel() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) Identifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) SetViewController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) TabState() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) Color() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) SetIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) ViewController() (*NSViewController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) TabView() (*NSTabView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) SetInitialFirstResponder() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) ToolTip() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) TabViewItemWithViewController() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewItem) DrawLabel() error {
  return fmt.Errorf("unimplemented")
}


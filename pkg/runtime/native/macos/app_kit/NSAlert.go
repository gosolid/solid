//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSAlert : objc.NSObject
*/

type NSAlert struct {
  *objc.NSObject

}

func (r *NSAlert) SetShowsHelp() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAlert) AccessoryView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAlert) Window() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAlert) BeginSheetModalForWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAlert) SetMessageText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAlert) SetInformativeText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAlert) Icon() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAlert) SetShowsSuppressionButton() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAlert) Buttons() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAlert) SetAlertStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAlert) SetAccessoryView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAlert) InformativeText() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAlert) Layout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAlert) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAlert) SuppressionButton() (*NSButton, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAlert) RunModal() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAlert) SetIcon() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAlert) HelpAnchor() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAlert) AlertWithError() (*NSAlert, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAlert) MessageText() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAlert) SetHelpAnchor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAlert) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAlert) AddButtonWithTitle() (*NSButton, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAlert) AlertStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAlert) ShowsHelp() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAlert) ShowsSuppressionButton() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


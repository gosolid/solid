//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSPrintPanel : objc.NSObject
*/

type NSPrintPanel struct {
  *objc.NSObject

}

func (r *NSPrintPanel) SetDefaultButtonTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) DefaultButtonTitle() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) PrintPanel() (*NSPrintPanel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) BeginSheetUsingPrintInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) RunModalWithPrintInfo() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) HelpAnchor() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) SetHelpAnchor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) JobStyleHint() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) SetJobStyleHint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) RemoveAccessoryController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) Options() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) SetOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) AddAccessoryController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) RunModal() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) AccessoryControllers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) PrintInfo() (*NSPrintInfo, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintPanel) BeginSheetWithPrintInfo() error {
  return fmt.Errorf("unimplemented")
}


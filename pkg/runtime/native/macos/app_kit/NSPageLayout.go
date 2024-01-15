//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSPageLayout : objc.NSObject
*/

type NSPageLayout struct {
  *objc.NSObject

}

func (r *NSPageLayout) PageLayout() (*NSPageLayout, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPageLayout) AddAccessoryController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPageLayout) RunModalWithPrintInfo() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPageLayout) RunModal() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPageLayout) AccessoryControllers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPageLayout) PrintInfo() (*NSPrintInfo, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPageLayout) RemoveAccessoryController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPageLayout) BeginSheetUsingPrintInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPageLayout) BeginSheetWithPrintInfo() error {
  return fmt.Errorf("unimplemented")
}


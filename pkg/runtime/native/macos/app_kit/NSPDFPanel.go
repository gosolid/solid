//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSPDFPanel : objc.NSObject
*/

type NSPDFPanel struct {
  *objc.NSObject

}

func (r *NSPDFPanel) BeginSheetWithPDFInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPDFPanel) AccessoryController() (*NSViewController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPDFPanel) SetAccessoryController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPDFPanel) Options() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPDFPanel) SetOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPDFPanel) DefaultFileName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPDFPanel) SetDefaultFileName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPDFPanel) Panel() (*NSPDFPanel, error) {
  return nil, fmt.Errorf("unimplemented")
}


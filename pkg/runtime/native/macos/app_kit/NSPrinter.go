//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSPrinter : objc.NSObject
*/

type NSPrinter struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSCoding
}

func (r *NSPrinter) PrinterWithName() (*NSPrinter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrinter) PrinterNames() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrinter) PrinterTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrinter) Type() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrinter) LanguageLevel() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPrinter) PrinterWithType() (*NSPrinter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrinter) PageSizeForPaper() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSPrinter) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrinter) DeviceDescription() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}


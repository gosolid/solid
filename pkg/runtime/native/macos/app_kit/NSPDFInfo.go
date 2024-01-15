//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSPDFInfo : objc.NSObject
*/

type NSPDFInfo struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSCoding
}

func (r *NSPDFInfo) TagNames() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPDFInfo) SetOrientation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPDFInfo) SetPaperSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPDFInfo) SetURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPDFInfo) SetFileExtensionHidden() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPDFInfo) SetTagNames() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPDFInfo) Orientation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPDFInfo) PaperSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSPDFInfo) Attributes() (*foundation.NSMutableDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPDFInfo) URL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPDFInfo) IsFileExtensionHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSPasteboard : objc.NSObject
*/

type NSPasteboard struct {
  *objc.NSObject

}

func (r *NSPasteboard) CanReadObjectForClasses() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) AddTypes() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) SetData() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) PropertyListForType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) Types() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) PasteboardWithName() (*NSPasteboard, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) PasteboardWithUniqueName() (*NSPasteboard, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) PrepareForNewContentsWithOptions() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) CanReadItemWithDataConformingToTypes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) AvailableTypeFromArray() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) ChangeCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) ReleaseGlobally() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) WriteObjects() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) IndexOfPasteboardItem() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) StringForType() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) PasteboardItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) DeclareTypes() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) SetPropertyList() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) DataForType() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) GeneralPasteboard() (*NSPasteboard, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) Name() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) ClearContents() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) ReadObjectsForClasses() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboard) SetString() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


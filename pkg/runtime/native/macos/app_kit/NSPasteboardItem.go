//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSPasteboardItem : objc.NSObject
*/

type NSPasteboardItem struct {
  *objc.NSObject
  *NSPasteboardWriting
  *NSPasteboardReading
}

func (r *NSPasteboardItem) SetString() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPasteboardItem) PropertyListForType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboardItem) AvailableTypeFromArray() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboardItem) SetDataProvider() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPasteboardItem) SetData() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPasteboardItem) SetPropertyList() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPasteboardItem) DataForType() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboardItem) StringForType() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboardItem) Types() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


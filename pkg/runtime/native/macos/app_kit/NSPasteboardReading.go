//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSPasteboardReading
*/

type NSPasteboardReading struct {

}

func (r *NSPasteboardReading) ReadableTypesForPasteboard() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboardReading) ReadingOptionsForType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPasteboardReading) InitWithPasteboardPropertyList() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


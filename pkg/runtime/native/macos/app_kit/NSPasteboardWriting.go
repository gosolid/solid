//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSPasteboardWriting
*/

type NSPasteboardWriting struct {

}

func (r *NSPasteboardWriting) WritableTypesForPasteboard() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPasteboardWriting) WritingOptionsForType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPasteboardWriting) PasteboardPropertyListForType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


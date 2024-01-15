//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSPasteboardItemDataProvider
*/

type NSPasteboardItemDataProvider struct {

}

func (r *NSPasteboardItemDataProvider) Pasteboard() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPasteboardItemDataProvider) PasteboardFinishedWithDataProvider() error {
  return fmt.Errorf("unimplemented")
}


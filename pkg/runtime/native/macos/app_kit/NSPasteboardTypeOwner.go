//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSPasteboardTypeOwner
*/

type NSPasteboardTypeOwner struct {

}

func (r *NSPasteboardTypeOwner) Pasteboard() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPasteboardTypeOwner) PasteboardChangedOwner() error {
  return fmt.Errorf("unimplemented")
}


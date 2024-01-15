//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSToolbarItemValidation
*/

type NSToolbarItemValidation struct {

}

func (r *NSToolbarItemValidation) ValidateToolbarItem() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


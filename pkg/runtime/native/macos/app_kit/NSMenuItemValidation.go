//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSMenuItemValidation
*/

type NSMenuItemValidation struct {

}

func (r *NSMenuItemValidation) ValidateMenuItem() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


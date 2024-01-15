//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSUserInterfaceValidations
*/

type NSUserInterfaceValidations struct {

}

func (r *NSUserInterfaceValidations) ValidateUserInterfaceItem() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


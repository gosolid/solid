//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSServicesMenuRequestor
*/

type NSServicesMenuRequestor struct {

}

func (r *NSServicesMenuRequestor) WriteSelectionToPasteboard() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSServicesMenuRequestor) ReadSelectionFromPasteboard() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


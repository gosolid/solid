//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSAlertDelegate
*/

type NSAlertDelegate struct {

}

func (r *NSAlertDelegate) AlertShowHelp() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


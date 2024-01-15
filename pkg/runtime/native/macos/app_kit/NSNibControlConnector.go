//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSNibControlConnector : AppKit.NSNibConnector
*/

type NSNibControlConnector struct {
  *NSNibConnector

}

func (r *NSNibControlConnector) EstablishConnection() error {
  return fmt.Errorf("unimplemented")
}


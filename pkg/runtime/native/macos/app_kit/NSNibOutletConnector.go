//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSNibOutletConnector : AppKit.NSNibConnector
*/

type NSNibOutletConnector struct {
  *NSNibConnector

}

func (r *NSNibOutletConnector) EstablishConnection() error {
  return fmt.Errorf("unimplemented")
}


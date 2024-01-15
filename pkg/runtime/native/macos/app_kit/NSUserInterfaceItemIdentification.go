//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSUserInterfaceItemIdentification
*/

type NSUserInterfaceItemIdentification struct {

}

func (r *NSUserInterfaceItemIdentification) Identifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceItemIdentification) SetIdentifier() error {
  return fmt.Errorf("unimplemented")
}


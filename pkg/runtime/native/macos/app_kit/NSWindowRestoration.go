//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSWindowRestoration
*/

type NSWindowRestoration struct {

}

func (r *NSWindowRestoration) RestoreWindowWithIdentifier() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSATSTypesetter : AppKit.NSTypesetter
*/

type NSATSTypesetter struct {
  *NSTypesetter

}

func (r *NSATSTypesetter) SharedTypesetter() (*NSATSTypesetter, error) {
  return nil, fmt.Errorf("unimplemented")
}


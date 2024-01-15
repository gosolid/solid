//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSIgnoreMisspelledWords
*/

type NSIgnoreMisspelledWords struct {

}

func (r *NSIgnoreMisspelledWords) IgnoreSpelling() error {
  return fmt.Errorf("unimplemented")
}


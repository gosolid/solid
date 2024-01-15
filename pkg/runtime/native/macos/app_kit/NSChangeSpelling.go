//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSChangeSpelling
*/

type NSChangeSpelling struct {

}

func (r *NSChangeSpelling) ChangeSpelling() error {
  return fmt.Errorf("unimplemented")
}


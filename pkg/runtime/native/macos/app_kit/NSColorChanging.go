//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSColorChanging
*/

type NSColorChanging struct {

}

func (r *NSColorChanging) ChangeColor() error {
  return fmt.Errorf("unimplemented")
}


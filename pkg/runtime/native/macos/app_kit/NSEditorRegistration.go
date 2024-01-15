//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSEditorRegistration
*/

type NSEditorRegistration struct {

}

func (r *NSEditorRegistration) ObjectDidBeginEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEditorRegistration) ObjectDidEndEditing() error {
  return fmt.Errorf("unimplemented")
}


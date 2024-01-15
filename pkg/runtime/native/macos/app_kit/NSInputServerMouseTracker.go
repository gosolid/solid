//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSInputServerMouseTracker
*/

type NSInputServerMouseTracker struct {

}

func (r *NSInputServerMouseTracker) MouseUpOnCharacterIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSInputServerMouseTracker) MouseDownOnCharacterIndex() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSInputServerMouseTracker) MouseDraggedOnCharacterIndex() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


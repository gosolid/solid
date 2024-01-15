//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSDraggingSource
*/

type NSDraggingSource struct {

}

func (r *NSDraggingSource) DraggingSession() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDraggingSource) IgnoreModifierKeysForDraggingSession() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


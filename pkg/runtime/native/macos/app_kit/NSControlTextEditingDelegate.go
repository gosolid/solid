//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSControlTextEditingDelegate
*/

type NSControlTextEditingDelegate struct {

}

func (r *NSControlTextEditingDelegate) ControlTextDidChange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControlTextEditingDelegate) Control() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSControlTextEditingDelegate) ControlTextDidBeginEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSControlTextEditingDelegate) ControlTextDidEndEditing() error {
  return fmt.Errorf("unimplemented")
}


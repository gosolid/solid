//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTextDelegate
*/

type NSTextDelegate struct {

}

func (r *NSTextDelegate) TextShouldEndEditing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextDelegate) TextDidBeginEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextDelegate) TextDidEndEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextDelegate) TextDidChange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextDelegate) TextShouldBeginEditing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


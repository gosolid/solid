//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSEditor
*/

type NSEditor struct {

}

func (r *NSEditor) DiscardEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEditor) CommitEditing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEditor) CommitEditingWithDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEditor) CommitEditingAndReturnError() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


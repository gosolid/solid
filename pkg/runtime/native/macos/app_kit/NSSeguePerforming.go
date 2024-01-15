//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSSeguePerforming
*/

type NSSeguePerforming struct {

}

func (r *NSSeguePerforming) PrepareForSegue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSeguePerforming) PerformSegueWithIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSeguePerforming) ShouldPerformSegueWithIdentifier() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


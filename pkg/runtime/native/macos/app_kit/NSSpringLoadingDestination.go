//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSSpringLoadingDestination
*/

type NSSpringLoadingDestination struct {

}

func (r *NSSpringLoadingDestination) SpringLoadingUpdated() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSpringLoadingDestination) SpringLoadingExited() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpringLoadingDestination) DraggingEnded() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpringLoadingDestination) SpringLoadingActivated() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpringLoadingDestination) SpringLoadingHighlightChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSpringLoadingDestination) SpringLoadingEntered() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


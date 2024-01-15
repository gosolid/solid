//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSDraggingDestination
*/

type NSDraggingDestination struct {

}

func (r *NSDraggingDestination) ConcludeDragOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingDestination) WantsPeriodicDraggingUpdates() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDraggingDestination) UpdateDraggingItemsForDrag() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingDestination) DraggingEntered() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDraggingDestination) DraggingUpdated() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDraggingDestination) DraggingExited() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingDestination) PrepareForDragOperation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDraggingDestination) PerformDragOperation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDraggingDestination) DraggingEnded() error {
  return fmt.Errorf("unimplemented")
}


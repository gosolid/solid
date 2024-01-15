//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSDraggingInfo
*/

type NSDraggingInfo struct {

}

func (r *NSDraggingInfo) DraggingSource() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) DraggingSequenceNumber() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) DraggingFormation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) SetDraggingFormation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) ResetSpringLoading() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) DraggingLocation() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) DraggedImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) DraggingPasteboard() (*NSPasteboard, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) NamesOfPromisedFilesDroppedAtDestination() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) DraggingDestinationWindow() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) SetNumberOfValidItemsForDrop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) SpringLoadingHighlight() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) EnumerateDraggingItemsWithOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) DraggedImageLocation() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) AnimatesToDestination() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) SetAnimatesToDestination() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) NumberOfValidItemsForDrop() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) SlideDraggedImageTo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingInfo) DraggingSourceOperationMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


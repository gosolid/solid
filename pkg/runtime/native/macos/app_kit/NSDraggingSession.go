//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSDraggingSession : objc.NSObject
*/

type NSDraggingSession struct {
  *objc.NSObject

}

func (r *NSDraggingSession) EnumerateDraggingItemsWithOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingSession) SetDraggingFormation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingSession) AnimatesToStartingPositionsOnCancelOrFail() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDraggingSession) DraggingLocation() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSDraggingSession) DraggingFormation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDraggingSession) SetAnimatesToStartingPositionsOnCancelOrFail() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingSession) DraggingLeaderIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDraggingSession) SetDraggingLeaderIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingSession) DraggingPasteboard() (*NSPasteboard, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingSession) DraggingSequenceNumber() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}


//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSResponder : objc.NSObject
*/

type NSResponder struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSResponder) MouseUp() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) RightMouseDragged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) KeyUp() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) BecomeFirstResponder() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSResponder) ShouldBeTreatedAsInkEvent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSResponder) QuickLookWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) AcceptsFirstResponder() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSResponder) OtherMouseDown() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) RotateWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) SwipeWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) SmartMagnifyWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) TouchesEndedWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) ScrollWheel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) TabletPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) BeginGestureWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) ResignFirstResponder() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSResponder) HelpRequested() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) NextResponder() (*NSResponder, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSResponder) MouseMoved() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) ChangeModeWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) TouchesBeganWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) SetNextResponder() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) ValidRequestorForSendType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSResponder) MouseDragged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) MouseExited() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) SetMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) OtherMouseDragged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) EndGestureWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) PressureChangeWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) FlushBufferedKeyEvents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) Menu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSResponder) TouchesMovedWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) TouchesCancelledWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) SupplementalTargetForAction() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSResponder) CursorUpdate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) TryToPerform() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSResponder) PerformKeyEquivalent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSResponder) NoResponderFor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) WantsForwardedScrollEventsForAxis() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSResponder) RightMouseUp() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) MouseEntered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) KeyDown() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSResponder) RightMouseDown() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) TabletProximity() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) ShowContextHelp() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) WantsScrollEventsForSwipeTrackingOnAxis() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSResponder) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSResponder) MouseDown() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) FlagsChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) MagnifyWithEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) OtherMouseUp() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSResponder) InterpretKeyEvents() error {
  return fmt.Errorf("unimplemented")
}


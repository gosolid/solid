//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSEvent : objc.NSObject
*/

type NSEvent struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSCoding
}

func (r *NSEvent) ButtonNumber() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) EventNumber() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) Data1() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) UniqueID() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) KeyRepeatInterval() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) Window() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) ModifierFlags() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) HasPreciseScrollingDeltas() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEvent) AbsoluteX() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) ButtonMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) IsEnteringProximity() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEvent) CharactersByApplyingModifiers() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) StopPeriodicEvents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEvent) MomentumPhase() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) Characters() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) KeyCode() (uint16, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) VendorID() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) EventWithCGEvent() (*NSEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) EventRef() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEvent) VendorPointingDeviceType() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) Phase() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) TrackSwipeEventWithOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEvent) MouseEventWithType() (*NSEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) DeltaY() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) Tilt() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSEvent) PressureBehavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) PressedMouseButtons() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) CoalescedTouchesForTouch() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) EnterExitEventWithType() (*NSEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) RemoveMonitor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEvent) WindowNumber() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) DeltaX() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) ScrollingDeltaY() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) TrackingNumber() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) DeviceID() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) EventWithEventRef() (*NSEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) VendorDefined() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) LocationInWindow() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSEvent) DeltaZ() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) CGEvent() (*core_graphics.CGEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) TangentialPressure() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) PointingDeviceID() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) MouseLocation() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSEvent) DoubleClickInterval() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) AddLocalMonitorForEventsMatchingMask() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) ClickCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) Stage() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) StageTransition() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) Timestamp() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) KeyEventWithType() (*NSEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) OtherEventWithType() (*NSEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) IsDirectionInvertedFromDevice() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEvent) IsARepeat() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEvent) Subtype() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) Rotation() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) PointingDeviceSerialNumber() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) AllTouches() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) Magnification() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) Data2() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) AbsoluteZ() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) SystemTabletID() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) TouchesForView() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) StartPeriodicEventsAfterDelay() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEvent) CapabilityMask() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) TouchesMatchingPhase() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) IsMouseCoalescingEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEvent) AssociatedEventsMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) Context() (*NSGraphicsContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) CharactersIgnoringModifiers() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) SetMouseCoalescingEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEvent) AbsoluteY() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) TabletID() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) PointingDeviceType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) KeyRepeatDelay() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) AddGlobalMonitorForEventsMatchingMask() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) Pressure() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) ScrollingDeltaX() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) UserData() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEvent) TrackingArea() (*NSTrackingArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEvent) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEvent) IsSwipeTrackingFromScrollEventsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSCursor : objc.NSObject
*/

type NSCursor struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *NSCursor) PointingHandCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) IBeamCursorForVerticalLayout() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) Unhide() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCursor) CurrentCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) ArrowCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) ResizeLeftRightCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) DisappearingItemCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) Hide() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCursor) ResizeRightCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) OpenHandCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) CrosshairCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) OperationNotAllowedCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) DragLinkCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) ResizeUpCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) ResizeDownCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) ResizeUpDownCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) SetHiddenUntilMouseMoves() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCursor) Pop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCursor) Push() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCursor) Set() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCursor) ResizeLeftCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) DragCopyCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) CurrentSystemCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) IBeamCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) ContextualMenuCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) InitWithImage() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) ClosedHandCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCursor) HotSpot() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}


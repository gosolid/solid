//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSClipView : AppKit.NSView
*/

type NSClipView struct {
  *NSView

}

func (r *NSClipView) ScrollToPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSClipView) DocumentCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSClipView) DocumentVisibleRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSClipView) SetAutomaticallyAdjustsContentInsets() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSClipView) SetContentInsets() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSClipView) ViewFrameChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSClipView) ConstrainBoundsRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSClipView) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSClipView) SetDrawsBackground() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSClipView) SetDocumentView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSClipView) DrawsBackground() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSClipView) DocumentView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSClipView) ContentInsets() (foundation.NSEdgeInsets, error) {
  return foundation.NSEdgeInsets{}, fmt.Errorf("unimplemented")
}

func (r *NSClipView) AutomaticallyAdjustsContentInsets() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSClipView) ViewBoundsChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSClipView) Autoscroll() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSClipView) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSClipView) DocumentRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSClipView) SetDocumentCursor() error {
  return fmt.Errorf("unimplemented")
}


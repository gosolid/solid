//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSScrollView : AppKit.NSView
*/

type NSScrollView struct {
  *NSView
  *NSTextFinderBarContainer
}

func (r *NSScrollView) SetAutohidesScrollers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetLineScroll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetHorizontalPageScroll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetUsesPredominantAxisScrolling() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) VerticalScroller() (*NSScroller, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetHorizontalLineScroll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) HorizontalPageScroll() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) HorizontalScrollElasticity() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) HasVerticalScroller() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetHasHorizontalScroller() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) PageScroll() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) MinMagnification() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) ContentSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) ContentView() (*NSClipView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) HorizontalScroller() (*NSScroller, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetAllowsMagnification() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetScrollerStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetContentInsets() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) Tile() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) ReflectScrolledClipView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) DocumentView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetContentView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) FlashScrollers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetHorizontalScroller() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) HorizontalLineScroll() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) AutohidesScrollers() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetVerticalLineScroll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) AutomaticallyAdjustsContentInsets() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) AddFloatingSubview() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) ContentInsets() (foundation.NSEdgeInsets, error) {
  return foundation.NSEdgeInsets{}, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) VerticalScrollElasticity() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) ScrollerInsets() (foundation.NSEdgeInsets, error) {
  return foundation.NSEdgeInsets{}, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) InitWithFrame() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) ScrollsDynamically() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) ScrollerStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) ScrollerKnobStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetVerticalPageScroll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) AllowsMagnification() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) MaxMagnification() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetMinMagnification() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) FrameSizeForContentSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) ScrollWheel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetPageScroll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) Magnification() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetDocumentView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) BorderType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) VerticalLineScroll() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) VerticalPageScroll() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetVerticalScroller() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetAutomaticallyAdjustsContentInsets() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetScrollerInsets() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) MagnifyToFitRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) DocumentVisibleRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) DrawsBackground() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetDrawsBackground() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetScrollsDynamically() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetVerticalScrollElasticity() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetMaxMagnification() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetMagnification() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) DocumentCursor() (*NSCursor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetBorderType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetHorizontalScrollElasticity() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) LineScroll() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetScrollerKnobStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) UsesPredominantAxisScrolling() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) ContentSizeForFrameSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetDocumentCursor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) SetHasVerticalScroller() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrollView) HasHorizontalScroller() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


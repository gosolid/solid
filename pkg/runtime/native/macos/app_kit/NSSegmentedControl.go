//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSSegmentedControl : AppKit.NSControl
*/

type NSSegmentedControl struct {
  *NSControl
  *NSUserInterfaceCompression
}

func (r *NSSegmentedControl) SetSegmentDistribution() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) MenuForSegment() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) IsEnabledForSegment() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) AlignmentForSegment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetSegmentStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetSelectedSegmentBezelColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetSpringLoaded() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) TrackingMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetToolTip() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetShowsMenuIndicator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SelectedSegment() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) IsSpringLoaded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SegmentStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) WidthForSegment() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetImageScaling() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetSegmentCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetTrackingMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SegmentDistribution() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) ImageScalingForSegment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) ShowsMenuIndicatorForSegment() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) MinimumSizeWithPrioritizedCompressionOptions() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SegmentCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SelectedSegmentBezelColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) ImageForSegment() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetSelected() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) IsSelectedForSegment() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) CompressWithPrioritizedCompressionOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) DoubleValueForSelectedSegment() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) IndexOfSelectedItem() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SelectSegmentWithTag() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) ToolTipForSegment() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) TagForSegment() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) LabelForSegment() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetTag() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) SetSelectedSegment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedControl) ActiveCompressionOptions() (*NSUserInterfaceCompressionOptions, error) {
  return nil, fmt.Errorf("unimplemented")
}


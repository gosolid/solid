//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSSegmentedCell : AppKit.NSActionCell
*/

type NSSegmentedCell struct {
  *NSActionCell

}

func (r *NSSegmentedCell) SelectSegmentWithTag() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) ImageScalingForSegment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) LabelForSegment() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) IsSelectedForSegment() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) ToolTipForSegment() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SetTrackingMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SegmentStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) MakePreviousSegmentKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) ImageForSegment() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SetMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) MenuForSegment() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) TagForSegment() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SetSelected() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SetToolTip() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) TrackingMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SetImageScaling() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SelectedSegment() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) WidthForSegment() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) DrawSegment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SetSegmentStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) MakeNextSegmentKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) IsEnabledForSegment() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SetTag() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SegmentCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SetSegmentCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSegmentedCell) SetSelectedSegment() error {
  return fmt.Errorf("unimplemented")
}


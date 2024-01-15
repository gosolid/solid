//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSLevelIndicator : AppKit.NSControl
*/

type NSLevelIndicator struct {
  *NSControl

}

func (r *NSLevelIndicator) WarningValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetNumberOfMajorTickMarks() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) PlaceholderVisibility() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetRatingPlaceholderImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) WarningFillColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetWarningFillColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) DrawsTieredCapacityLevels() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetDrawsTieredCapacityLevels() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) RatingPlaceholderImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) MaxValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) TickMarkPosition() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetRatingImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) RatingImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) MinValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetTickMarkPosition() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) NumberOfTickMarks() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) FillColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) CriticalFillColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetPlaceholderVisibility() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) IsEditable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetWarningValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetCriticalValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetNumberOfTickMarks() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetFillColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) TickMarkValueAtIndex() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetMinValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) NumberOfMajorTickMarks() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) RectOfTickMarkAtIndex() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetLevelIndicatorStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetCriticalFillColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) LevelIndicatorStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetEditable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) SetMaxValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicator) CriticalValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}


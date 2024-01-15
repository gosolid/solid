//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSLevelIndicatorCell : AppKit.NSActionCell
*/

type NSLevelIndicatorCell struct {
  *NSActionCell

}

func (r *NSLevelIndicatorCell) InitWithLevelIndicatorStyle() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) RectOfTickMarkAtIndex() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) TickMarkValueAtIndex() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) LevelIndicatorStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) SetCriticalValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) MinValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) SetMinValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) MaxValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) SetMaxValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) SetWarningValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) SetNumberOfTickMarks() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) NumberOfMajorTickMarks() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) SetLevelIndicatorStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) WarningValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) CriticalValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) SetTickMarkPosition() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) TickMarkPosition() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) NumberOfTickMarks() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLevelIndicatorCell) SetNumberOfMajorTickMarks() error {
  return fmt.Errorf("unimplemented")
}


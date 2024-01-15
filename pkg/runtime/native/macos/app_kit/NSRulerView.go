//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSRulerView : AppKit.NSView
*/

type NSRulerView struct {
  *NSView

}

func (r *NSRulerView) InvalidateHashMarks() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) DrawMarkersInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) SetReservedThicknessForAccessoryView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) InitWithScrollView() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) DrawHashMarksAndLabelsInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) ReservedThicknessForAccessoryView() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) ClientView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) MoveRulerlineFromLocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) OriginOffset() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) SetMarkers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) MeasurementUnits() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) IsFlipped() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) AddMarker() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) ScrollView() (*NSScrollView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) RequiredThickness() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) RuleThickness() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) RegisterUnitWithName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) SetScrollView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) SetMeasurementUnits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) AccessoryView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) ReservedThicknessForMarkers() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) Markers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) SetAccessoryView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) Orientation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) BaselineLocation() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) SetRuleThickness() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) SetOriginOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) SetClientView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) RemoveMarker() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) TrackMarker() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRulerView) SetOrientation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerView) SetReservedThicknessForMarkers() error {
  return fmt.Errorf("unimplemented")
}


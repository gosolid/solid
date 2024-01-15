//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSDateComponentsFormatter : Foundation.NSFormatter
*/

type NSDateComponentsFormatter struct {
  *NSFormatter

}

func (r *NSDateComponentsFormatter) Calendar() (*NSCalendar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) AllowsFractionalUnits() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) SetAllowsFractionalUnits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) IncludesApproximationPhrase() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) SetIncludesApproximationPhrase() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) FormattingContext() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) StringForObjectValue() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) SetUnitsStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) SetCollapsesLargestUnit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) SetIncludesTimeRemainingPhrase() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) StringFromDate() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) SetZeroFormattingBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) SetAllowedUnits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) StringFromTimeInterval() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) ReferenceDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) StringFromDateComponents() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) ZeroFormattingBehavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) SetMaximumUnitCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) SetFormattingContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) GetObjectValue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) UnitsStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) LocalizedStringFromDateComponents() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) CollapsesLargestUnit() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) SetReferenceDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) MaximumUnitCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) IncludesTimeRemainingPhrase() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) AllowedUnits() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponentsFormatter) SetCalendar() error {
  return fmt.Errorf("unimplemented")
}


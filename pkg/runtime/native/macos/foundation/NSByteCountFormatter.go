//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSByteCountFormatter : Foundation.NSFormatter
*/

type NSByteCountFormatter struct {
  *NSFormatter

}

func (r *NSByteCountFormatter) SetFormattingContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) StringFromMeasurement() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) StringForObjectValue() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetIncludesActualByteCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) FormattingContext() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) StringFromByteCount() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) CountStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetAllowsNonnumericFormatting() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) IncludesUnit() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetIncludesUnit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) IncludesCount() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetIncludesCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) AllowedUnits() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetAllowedUnits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetCountStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) AllowsNonnumericFormatting() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetZeroPadsFractionDigits() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) IncludesActualByteCount() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) IsAdaptive() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetAdaptive() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) ZeroPadsFractionDigits() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


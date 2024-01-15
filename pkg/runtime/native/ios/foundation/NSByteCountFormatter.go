//js:package native/ios/foundation
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

func (r *NSByteCountFormatter) SetCountStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) AllowsNonnumericFormatting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetAllowsNonnumericFormatting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) ZeroPadsFractionDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetZeroPadsFractionDigits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetFormattingContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) AllowedUnits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetIncludesUnit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) IncludesActualByteCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) IsAdaptive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) IncludesCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetIncludesCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetAdaptive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) FormattingContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) StringFromMeasurement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetAllowedUnits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) CountStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) IncludesUnit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) StringFromByteCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) StringForObjectValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSByteCountFormatter) SetIncludesActualByteCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


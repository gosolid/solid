//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSLengthFormatter : Foundation.NSFormatter
*/

type NSLengthFormatter struct {
  *NSFormatter

}

func (r *NSLengthFormatter) UnitStringFromValue() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) GetObjectValue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) UnitStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) IsForPersonHeightUse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) SetForPersonHeightUse() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) StringFromValue() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) StringFromMeters() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) UnitStringFromMeters() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) NumberFormatter() (*NSNumberFormatter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) SetNumberFormatter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) SetUnitStyle() error {
  return fmt.Errorf("unimplemented")
}


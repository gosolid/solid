//js:package native/ios/foundation
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

func (r *NSLengthFormatter) StringFromValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) UnitStringFromMeters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) NumberFormatter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) IsForPersonHeightUse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) SetForPersonHeightUse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) StringFromMeters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) UnitStringFromValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) GetObjectValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) SetNumberFormatter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) UnitStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLengthFormatter) SetUnitStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


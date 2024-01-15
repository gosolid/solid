//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMassFormatter : Foundation.NSFormatter
*/

type NSMassFormatter struct {
  *NSFormatter

}

func (r *NSMassFormatter) StringFromKilograms() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) UnitStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) SetUnitStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) IsForPersonMassUse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) StringFromValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) UnitStringFromKilograms() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) GetObjectValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) NumberFormatter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) SetNumberFormatter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) SetForPersonMassUse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMassFormatter) UnitStringFromValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


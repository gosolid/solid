//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSEnergyFormatter : Foundation.NSFormatter
*/

type NSEnergyFormatter struct {
  *NSFormatter

}

func (r *NSEnergyFormatter) UnitStringFromValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) UnitStringFromJoules() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) GetObjectValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) SetForFoodEnergyUse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) StringFromValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) StringFromJoules() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) NumberFormatter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) SetNumberFormatter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) UnitStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) SetUnitStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) IsForFoodEnergyUse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


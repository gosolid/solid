//js:package native/macos/foundation
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

func (r *NSEnergyFormatter) SetNumberFormatter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) SetUnitStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) IsForFoodEnergyUse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) StringFromValue() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) StringFromJoules() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) UnitStringFromJoules() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) GetObjectValue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) UnitStringFromValue() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) NumberFormatter() (*NSNumberFormatter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) UnitStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEnergyFormatter) SetForFoodEnergyUse() error {
  return fmt.Errorf("unimplemented")
}


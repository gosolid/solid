//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitElectricPotentialDifference : Foundation.NSDimension
*/

type NSUnitElectricPotentialDifference struct {
  *NSDimension
  *NSSecureCoding
}

func (r *NSUnitElectricPotentialDifference) Kilovolts() (*NSUnitElectricPotentialDifference, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricPotentialDifference) Volts() (*NSUnitElectricPotentialDifference, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricPotentialDifference) Millivolts() (*NSUnitElectricPotentialDifference, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricPotentialDifference) Microvolts() (*NSUnitElectricPotentialDifference, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricPotentialDifference) Megavolts() (*NSUnitElectricPotentialDifference, error) {
  return nil, fmt.Errorf("unimplemented")
}


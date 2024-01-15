//js:package native/ios/foundation
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

}

func (r *NSUnitElectricPotentialDifference) Microvolts() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricPotentialDifference) Megavolts() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricPotentialDifference) Kilovolts() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricPotentialDifference) Volts() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricPotentialDifference) Millivolts() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


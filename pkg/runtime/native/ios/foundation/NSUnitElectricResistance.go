//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitElectricResistance : Foundation.NSDimension
*/

type NSUnitElectricResistance struct {
  *NSDimension

}

func (r *NSUnitElectricResistance) Milliohms() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricResistance) Microohms() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricResistance) Megaohms() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricResistance) Kiloohms() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricResistance) Ohms() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


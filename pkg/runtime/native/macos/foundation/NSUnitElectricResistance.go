//js:package native/macos/foundation
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
  *NSSecureCoding
}

func (r *NSUnitElectricResistance) Kiloohms() (*NSUnitElectricResistance, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricResistance) Ohms() (*NSUnitElectricResistance, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricResistance) Milliohms() (*NSUnitElectricResistance, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricResistance) Microohms() (*NSUnitElectricResistance, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricResistance) Megaohms() (*NSUnitElectricResistance, error) {
  return nil, fmt.Errorf("unimplemented")
}


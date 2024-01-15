//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitElectricCharge : Foundation.NSDimension
*/

type NSUnitElectricCharge struct {
  *NSDimension
  *NSSecureCoding
}

func (r *NSUnitElectricCharge) MicroampereHours() (*NSUnitElectricCharge, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCharge) Coulombs() (*NSUnitElectricCharge, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCharge) MegaampereHours() (*NSUnitElectricCharge, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCharge) KiloampereHours() (*NSUnitElectricCharge, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCharge) AmpereHours() (*NSUnitElectricCharge, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCharge) MilliampereHours() (*NSUnitElectricCharge, error) {
  return nil, fmt.Errorf("unimplemented")
}


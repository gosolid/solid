//js:package native/ios/foundation
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

}

func (r *NSUnitElectricCharge) AmpereHours() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCharge) MilliampereHours() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCharge) MicroampereHours() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCharge) Coulombs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCharge) MegaampereHours() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCharge) KiloampereHours() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


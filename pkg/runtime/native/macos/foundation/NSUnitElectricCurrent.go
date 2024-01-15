//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitElectricCurrent : Foundation.NSDimension
*/

type NSUnitElectricCurrent struct {
  *NSDimension
  *NSSecureCoding
}

func (r *NSUnitElectricCurrent) Megaamperes() (*NSUnitElectricCurrent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCurrent) Kiloamperes() (*NSUnitElectricCurrent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCurrent) Amperes() (*NSUnitElectricCurrent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCurrent) Milliamperes() (*NSUnitElectricCurrent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCurrent) Microamperes() (*NSUnitElectricCurrent, error) {
  return nil, fmt.Errorf("unimplemented")
}


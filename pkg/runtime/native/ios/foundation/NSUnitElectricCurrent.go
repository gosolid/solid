//js:package native/ios/foundation
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

}

func (r *NSUnitElectricCurrent) Amperes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCurrent) Milliamperes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCurrent) Microamperes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCurrent) Megaamperes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitElectricCurrent) Kiloamperes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


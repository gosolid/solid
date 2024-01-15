//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitEnergy : Foundation.NSDimension
*/

type NSUnitEnergy struct {
  *NSDimension
  *NSSecureCoding
}

func (r *NSUnitEnergy) Kilojoules() (*NSUnitEnergy, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitEnergy) Joules() (*NSUnitEnergy, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitEnergy) Kilocalories() (*NSUnitEnergy, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitEnergy) Calories() (*NSUnitEnergy, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitEnergy) KilowattHours() (*NSUnitEnergy, error) {
  return nil, fmt.Errorf("unimplemented")
}


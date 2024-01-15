//js:package native/ios/foundation
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

}

func (r *NSUnitEnergy) Kilojoules() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitEnergy) Joules() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitEnergy) Kilocalories() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitEnergy) Calories() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitEnergy) KilowattHours() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


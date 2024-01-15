//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitTemperature : Foundation.NSDimension
*/

type NSUnitTemperature struct {
  *NSDimension

}

func (r *NSUnitTemperature) Kelvin() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitTemperature) Celsius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitTemperature) Fahrenheit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


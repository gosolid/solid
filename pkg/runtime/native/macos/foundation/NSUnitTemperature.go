//js:package native/macos/foundation
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
  *NSSecureCoding
}

func (r *NSUnitTemperature) Kelvin() (*NSUnitTemperature, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitTemperature) Celsius() (*NSUnitTemperature, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitTemperature) Fahrenheit() (*NSUnitTemperature, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitFuelEfficiency : Foundation.NSDimension
*/

type NSUnitFuelEfficiency struct {
  *NSDimension
  *NSSecureCoding
}

func (r *NSUnitFuelEfficiency) MilesPerImperialGallon() (*NSUnitFuelEfficiency, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFuelEfficiency) MilesPerGallon() (*NSUnitFuelEfficiency, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFuelEfficiency) LitersPer100Kilometers() (*NSUnitFuelEfficiency, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/foundation
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

}

func (r *NSUnitFuelEfficiency) LitersPer100Kilometers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFuelEfficiency) MilesPerImperialGallon() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFuelEfficiency) MilesPerGallon() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


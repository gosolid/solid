//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitPressure : Foundation.NSDimension
*/

type NSUnitPressure struct {
  *NSDimension

}

func (r *NSUnitPressure) Megapascals() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) Kilopascals() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) Bars() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) Millibars() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) PoundsForcePerSquareInch() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) NewtonsPerMetersSquared() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) Gigapascals() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) MillimetersOfMercury() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) Hectopascals() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) InchesOfMercury() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


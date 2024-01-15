//js:package native/macos/foundation
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
  *NSSecureCoding
}

func (r *NSUnitPressure) Megapascals() (*NSUnitPressure, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) Kilopascals() (*NSUnitPressure, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) InchesOfMercury() (*NSUnitPressure, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) PoundsForcePerSquareInch() (*NSUnitPressure, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) MillimetersOfMercury() (*NSUnitPressure, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) NewtonsPerMetersSquared() (*NSUnitPressure, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) Gigapascals() (*NSUnitPressure, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) Hectopascals() (*NSUnitPressure, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) Bars() (*NSUnitPressure, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPressure) Millibars() (*NSUnitPressure, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitConcentrationMass : Foundation.NSDimension
*/

type NSUnitConcentrationMass struct {
  *NSDimension
  *NSSecureCoding
}

func (r *NSUnitConcentrationMass) MilligramsPerDeciliter() (*NSUnitConcentrationMass, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitConcentrationMass) MillimolesPerLiterWithGramsPerMole() (*NSUnitConcentrationMass, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitConcentrationMass) GramsPerLiter() (*NSUnitConcentrationMass, error) {
  return nil, fmt.Errorf("unimplemented")
}


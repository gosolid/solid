//js:package native/ios/foundation
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

}

func (r *NSUnitConcentrationMass) MillimolesPerLiterWithGramsPerMole() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitConcentrationMass) GramsPerLiter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitConcentrationMass) MilligramsPerDeciliter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


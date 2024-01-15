//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitConverterLinear : Foundation.NSUnitConverter
*/

type NSUnitConverterLinear struct {
  *NSUnitConverter
  *NSSecureCoding
}

func (r *NSUnitConverterLinear) InitWithCoefficient() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitConverterLinear) Coefficient() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSUnitConverterLinear) Constant() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}


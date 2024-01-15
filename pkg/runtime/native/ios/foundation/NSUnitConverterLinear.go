//js:package native/ios/foundation
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

}

func (r *NSUnitConverterLinear) InitWithCoefficient() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitConverterLinear) Coefficient() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitConverterLinear) Constant() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


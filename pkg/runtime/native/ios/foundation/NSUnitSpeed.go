//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitSpeed : Foundation.NSDimension
*/

type NSUnitSpeed struct {
  *NSDimension

}

func (r *NSUnitSpeed) KilometersPerHour() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitSpeed) MilesPerHour() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitSpeed) Knots() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitSpeed) MetersPerSecond() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


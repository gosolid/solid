//js:package native/macos/foundation
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
  *NSSecureCoding
}

func (r *NSUnitSpeed) MetersPerSecond() (*NSUnitSpeed, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitSpeed) KilometersPerHour() (*NSUnitSpeed, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitSpeed) MilesPerHour() (*NSUnitSpeed, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitSpeed) Knots() (*NSUnitSpeed, error) {
  return nil, fmt.Errorf("unimplemented")
}


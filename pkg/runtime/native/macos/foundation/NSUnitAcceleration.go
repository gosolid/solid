//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitAcceleration : Foundation.NSDimension
*/

type NSUnitAcceleration struct {
  *NSDimension
  *NSSecureCoding
}

func (r *NSUnitAcceleration) MetersPerSecondSquared() (*NSUnitAcceleration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitAcceleration) Gravity() (*NSUnitAcceleration, error) {
  return nil, fmt.Errorf("unimplemented")
}


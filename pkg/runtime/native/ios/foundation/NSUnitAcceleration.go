//js:package native/ios/foundation
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

}

func (r *NSUnitAcceleration) Gravity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitAcceleration) MetersPerSecondSquared() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


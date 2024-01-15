//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitAngle : Foundation.NSDimension
*/

type NSUnitAngle struct {
  *NSDimension
  *NSSecureCoding
}

func (r *NSUnitAngle) Degrees() (*NSUnitAngle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitAngle) ArcMinutes() (*NSUnitAngle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitAngle) ArcSeconds() (*NSUnitAngle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitAngle) Radians() (*NSUnitAngle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitAngle) Gradians() (*NSUnitAngle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitAngle) Revolutions() (*NSUnitAngle, error) {
  return nil, fmt.Errorf("unimplemented")
}


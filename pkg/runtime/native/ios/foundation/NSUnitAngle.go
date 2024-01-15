//js:package native/ios/foundation
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

}

func (r *NSUnitAngle) Degrees() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitAngle) ArcMinutes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitAngle) ArcSeconds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitAngle) Radians() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitAngle) Gradians() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitAngle) Revolutions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


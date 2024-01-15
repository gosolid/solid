//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface QuartzCore.CAConstraint : objc.NSObject
*/

type CAConstraint struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *CAConstraint) ConstraintWithAttribute() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAConstraint) InitWithAttribute() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAConstraint) Attribute() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAConstraint) SourceName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAConstraint) SourceAttribute() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAConstraint) Scale() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAConstraint) Offset() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}


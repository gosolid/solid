//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreLocation.CLCondition : objc.NSObject
*/

type CLCondition struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CLCondition) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLCondition) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


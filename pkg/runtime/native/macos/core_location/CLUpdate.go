//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CoreLocation.CLUpdate : objc.NSObject
*/

type CLUpdate struct {
  *objc.NSObject

}

func (r *CLUpdate) Location() (*CLLocation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLUpdate) IsStationary() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


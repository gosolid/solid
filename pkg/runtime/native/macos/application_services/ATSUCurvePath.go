//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
struct ApplicationServices.ATSUCurvePath
*/

type ATSUCurvePath struct {
  Vectors uint `v8:"vectors"`
  ControlBits [1]uint `v8:"controlBits"`
  Vector [1]core_foundation.CGPoint `v8:"vector"`
}

//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct ApplicationServices.FMInput
*/

type FMInput struct {
  Family int16 `v8:"family"`
  Size int16 `v8:"size"`
  Face byte `v8:"face"`
  NeedBits byte `v8:"needBits"`
  Device int16 `v8:"device"`
  Numer objc.Point `v8:"numer"`
  Denom objc.Point `v8:"denom"`
}

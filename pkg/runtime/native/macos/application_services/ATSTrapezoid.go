//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct ApplicationServices.ATSTrapezoid
*/

type ATSTrapezoid struct {
  UpperLeft objc.FixedPoint `v8:"upperLeft"`
  UpperRight objc.FixedPoint `v8:"upperRight"`
  LowerRight objc.FixedPoint `v8:"lowerRight"`
  LowerLeft objc.FixedPoint `v8:"lowerLeft"`
}

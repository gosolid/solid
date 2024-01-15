//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
struct ApplicationServices.ATSGlyphScreenMetrics
*/

type ATSGlyphScreenMetrics struct {
  DeviceAdvance core_foundation.CGPoint `v8:"deviceAdvance"`
  TopLeft core_foundation.CGPoint `v8:"topLeft"`
  Height uint `v8:"height"`
  Width uint `v8:"width"`
  SideBearing core_foundation.CGPoint `v8:"sideBearing"`
  OtherSideBearing core_foundation.CGPoint `v8:"otherSideBearing"`
}

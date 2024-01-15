//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
struct ApplicationServices.ATSGlyphIdealMetrics
*/

type ATSGlyphIdealMetrics struct {
  Advance core_foundation.CGPoint `v8:"advance"`
  SideBearing core_foundation.CGPoint `v8:"sideBearing"`
  OtherSideBearing core_foundation.CGPoint `v8:"otherSideBearing"`
}

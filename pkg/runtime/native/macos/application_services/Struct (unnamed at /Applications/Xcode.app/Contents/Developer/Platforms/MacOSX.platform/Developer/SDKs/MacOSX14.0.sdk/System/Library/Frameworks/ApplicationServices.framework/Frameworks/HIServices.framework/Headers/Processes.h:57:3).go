//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct ApplicationServices.struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/ApplicationServices.framework/Frameworks/HIServices.framework/Headers/Processes.h:57:3)
*/

type Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/ApplicationServices.framework/Frameworks/HIServices.framework/Headers/Processes.h:57:3) struct {
  What uint16 `v8:"what"`
  Message uint `v8:"message"`
  When uint `v8:"when"`
  Where objc.Point `v8:"where"`
  Modifiers uint16 `v8:"modifiers"`
}

//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.AppParameters
*/

type AppParameters struct {
  TheMsgEvent Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/ApplicationServices.framework/Frameworks/HIServices.framework/Headers/Processes.h:57:3) `v8:"theMsgEvent"`
  EventRefCon uint `v8:"eventRefCon"`
  MessageLength uint `v8:"messageLength"`
}

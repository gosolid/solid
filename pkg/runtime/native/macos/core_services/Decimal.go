//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.decimal
*/

type Decimal struct {
  Sgn byte `v8:"sgn"`
  Unused byte `v8:"unused"`
  Exp int16 `v8:"exp"`
  Sig Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/CoreServices.framework/Frameworks/CarbonCore.framework/Headers/fp.h:1320:5) `v8:"sig"`
}

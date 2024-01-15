//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Foundation.struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/Foundation.framework/Headers/NSScriptCommand.h:37:5)
*/

type Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/Foundation.framework/Headers/NSScriptCommand.h:37:5) struct {
  HasEvaluatedReceivers uint `v8:"hasEvaluatedReceivers"`
  HasEvaluatedArguments uint `v8:"hasEvaluatedArguments"`
  RESERVED uint `v8:"rESERVED"`
}

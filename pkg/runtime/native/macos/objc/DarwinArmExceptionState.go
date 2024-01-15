//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__darwin_arm_exception_state
*/

type DarwinArmExceptionState struct {
  Exception uint `v8:"exception"`
  Fsr uint `v8:"fsr"`
  Far uint `v8:"far"`
}

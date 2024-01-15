//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__darwin_arm_exception_state64
*/

type DarwinArmExceptionState64 struct {
  Far uint64 `v8:"far"`
  Esr uint `v8:"esr"`
  Exception uint `v8:"exception"`
}

//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__darwin_arm_thread_state
*/

type DarwinArmThreadState struct {
  R [13]uint `v8:"r"`
  Sp uint `v8:"sp"`
  Lr uint `v8:"lr"`
  Pc uint `v8:"pc"`
  Cpsr uint `v8:"cpsr"`
}

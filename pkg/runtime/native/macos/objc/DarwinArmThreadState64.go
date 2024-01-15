//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__darwin_arm_thread_state64
*/

type DarwinArmThreadState64 struct {
  X [29]uint64 `v8:"x"`
  Fp uint64 `v8:"fp"`
  Lr uint64 `v8:"lr"`
  Sp uint64 `v8:"sp"`
  Pc uint64 `v8:"pc"`
  Cpsr uint `v8:"cpsr"`
  Pad uint `v8:"pad"`
}

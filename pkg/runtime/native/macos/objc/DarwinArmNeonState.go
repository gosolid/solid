//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__darwin_arm_neon_state
*/

type DarwinArmNeonState struct {
  V [16]uint64 `v8:"v"`
  Fpsr uint `v8:"fpsr"`
  Fpcr uint `v8:"fpcr"`
}

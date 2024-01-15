//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__darwin_arm_vfp_state
*/

type DarwinArmVfpState struct {
  R [64]uint `v8:"r"`
  Fpscr uint `v8:"fpscr"`
}

//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__arm_legacy_debug_state
*/

type ArmLegacyDebugState struct {
  Bvr [16]uint `v8:"bvr"`
  Bcr [16]uint `v8:"bcr"`
  Wvr [16]uint `v8:"wvr"`
  Wcr [16]uint `v8:"wcr"`
}

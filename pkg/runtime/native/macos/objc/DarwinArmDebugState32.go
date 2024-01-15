//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__darwin_arm_debug_state32
*/

type DarwinArmDebugState32 struct {
  Bvr [16]uint `v8:"bvr"`
  Bcr [16]uint `v8:"bcr"`
  Wvr [16]uint `v8:"wvr"`
  Wcr [16]uint `v8:"wcr"`
  MdscrEl1 uint64 `v8:"mdscrEl1"`
}

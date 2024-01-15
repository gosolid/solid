//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__darwin_arm_debug_state64
*/

type DarwinArmDebugState64 struct {
  Bvr [16]uint64 `v8:"bvr"`
  Bcr [16]uint64 `v8:"bcr"`
  Wvr [16]uint64 `v8:"wvr"`
  Wcr [16]uint64 `v8:"wcr"`
  MdscrEl1 uint64 `v8:"mdscrEl1"`
}

//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__darwin_arm_cpmu_state64
*/

type DarwinArmCpmuState64 struct {
  Ctrs [16]uint64 `v8:"ctrs"`
}

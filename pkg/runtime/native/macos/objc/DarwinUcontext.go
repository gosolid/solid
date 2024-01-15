//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__darwin_ucontext
*/

type DarwinUcontext struct {
  UcOnstack int `v8:"ucOnstack"`
  UcSigmask uint `v8:"ucSigmask"`
  UcStack DarwinSigaltstack `v8:"ucStack"`
  UcLink DarwinUcontext `v8:"ucLink"`
  UcMcsize uint64 `v8:"ucMcsize"`
  UcMcontext DarwinMcontext64 `v8:"ucMcontext"`
}

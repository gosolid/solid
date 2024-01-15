//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__darwin_mcontext64
*/

type DarwinMcontext64 struct {
  Es DarwinArmExceptionState64 `v8:"es"`
  Ss DarwinArmThreadState64 `v8:"ss"`
  Ns DarwinArmNeonState64 `v8:"ns"`
}

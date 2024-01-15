//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__darwin_mcontext32
*/

type DarwinMcontext32 struct {
  Es DarwinArmExceptionState `v8:"es"`
  Ss DarwinArmThreadState `v8:"ss"`
  Fs DarwinArmVfpState `v8:"fs"`
}

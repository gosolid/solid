//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.rlimit
*/

type Rlimit struct {
  RlimCur uint64 `v8:"rlimCur"`
  RlimMax uint64 `v8:"rlimMax"`
}

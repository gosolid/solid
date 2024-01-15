//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__darwin_sigaltstack
*/

type DarwinSigaltstack struct {
  SsSp void `v8:"ssSp"`
  SsSize uint64 `v8:"ssSize"`
  SsFlags int `v8:"ssFlags"`
}

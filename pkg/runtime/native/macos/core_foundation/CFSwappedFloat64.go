//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFSwappedFloat64
*/

type CFSwappedFloat64 struct {
  V uint64 `v8:"v"`
}

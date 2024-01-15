//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFSwappedFloat32
*/

type CFSwappedFloat32 struct {
  V uint `v8:"v"`
}

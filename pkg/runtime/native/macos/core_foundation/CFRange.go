//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFRange
*/

type CFRange struct {
  Location int64 `v8:"location"`
  Length int64 `v8:"length"`
}

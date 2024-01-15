//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Foundation.NSSwappedDouble
*/

type NSSwappedDouble struct {
  V uint64 `v8:"v"`
}

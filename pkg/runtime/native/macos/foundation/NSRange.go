//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Foundation.NSRange
*/

type NSRange struct {
  Location uint64 `v8:"location"`
  Length uint64 `v8:"length"`
}

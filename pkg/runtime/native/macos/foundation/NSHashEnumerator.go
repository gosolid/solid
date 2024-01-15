//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Foundation.NSHashEnumerator
*/

type NSHashEnumerator struct {
  Pi uint64 `v8:"pi"`
  Si uint64 `v8:"si"`
  Bs void `v8:"bs"`
}

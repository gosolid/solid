//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Foundation.NSMapEnumerator
*/

type NSMapEnumerator struct {
  Pi uint64 `v8:"pi"`
  Si uint64 `v8:"si"`
  Bs void `v8:"bs"`
}

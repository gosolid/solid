//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.ldiv_t
*/

type LdivT struct {
  Quot int64 `v8:"quot"`
  Rem int64 `v8:"rem"`
}

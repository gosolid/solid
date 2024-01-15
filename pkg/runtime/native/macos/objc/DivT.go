//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.div_t
*/

type DivT struct {
  Quot int `v8:"quot"`
  Rem int `v8:"rem"`
}

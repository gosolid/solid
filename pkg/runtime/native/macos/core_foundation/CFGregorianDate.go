//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFGregorianDate
*/

type CFGregorianDate struct {
  Year int `v8:"year"`
  Month byte `v8:"month"`
  Day byte `v8:"day"`
  Hour byte `v8:"hour"`
  Minute byte `v8:"minute"`
  Second float64 `v8:"second"`
}

//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CGVector
*/

type CGVector struct {
  Dx float64 `v8:"dx"`
  Dy float64 `v8:"dy"`
}

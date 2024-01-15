//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CGPoint
*/

type CGPoint struct {
  X float64 `v8:"x"`
  Y float64 `v8:"y"`
}

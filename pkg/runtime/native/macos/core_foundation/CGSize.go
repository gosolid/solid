//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CGSize
*/

type CGSize struct {
  Width float64 `v8:"width"`
  Height float64 `v8:"height"`
}

//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFStreamError
*/

type CFStreamError struct {
  Domain int64 `v8:"domain"`
  Error int `v8:"error"`
}
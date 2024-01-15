//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DDMap
*/

type DDMap struct {
  DdBlock uint `v8:"ddBlock"`
  DdSize uint16 `v8:"ddSize"`
  DdType uint16 `v8:"ddType"`
}

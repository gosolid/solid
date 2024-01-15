//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOVirtualRange
*/

type IOVirtualRange struct {
  Address uint64 `v8:"address"`
  Length uint64 `v8:"length"`
}

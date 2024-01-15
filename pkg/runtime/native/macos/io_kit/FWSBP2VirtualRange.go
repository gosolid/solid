//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.FWSBP2VirtualRange
*/

type FWSBP2VirtualRange struct {
  Address void `v8:"address"`
  Length uint `v8:"length"`
}

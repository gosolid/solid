//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IORPCMessage
*/

type IORPCMessage struct {
  Msgid uint64 `v8:"msgid"`
  Flags uint64 `v8:"flags"`
  ObjectRefs uint64 `v8:"objectRefs"`
  Objects [0]uint64 `v8:"objects"`
}

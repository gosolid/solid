//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IODataQueueEntry
*/

type IODataQueueEntry struct {
  Size uint `v8:"size"`
  Data [4]byte `v8:"data"`
}

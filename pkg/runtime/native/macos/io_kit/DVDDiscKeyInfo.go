//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DVDDiscKeyInfo
*/

type DVDDiscKeyInfo struct {
  DataLength [2]byte `v8:"dataLength"`
  Reserved [2]byte `v8:"reserved"`
  DiscKeyStructures [2048]byte `v8:"discKeyStructures"`
}

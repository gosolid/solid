//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DVDKey2Info
*/

type DVDKey2Info struct {
  DataLength [2]byte `v8:"dataLength"`
  Reserved [2]byte `v8:"reserved"`
  Key2Value [5]byte `v8:"key2Value"`
  Reserved2 [3]byte `v8:"reserved2"`
}

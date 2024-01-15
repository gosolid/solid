//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DVDKey1Info
*/

type DVDKey1Info struct {
  DataLength [2]byte `v8:"dataLength"`
  Reserved [2]byte `v8:"reserved"`
  Key1Value [5]byte `v8:"key1Value"`
  Reserved2 [3]byte `v8:"reserved2"`
}

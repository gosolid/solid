//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DVDManufacturingInfo
*/

type DVDManufacturingInfo struct {
  DataLength [2]byte `v8:"dataLength"`
  Reserved [2]byte `v8:"reserved"`
  DiscManufacturingInfo [2048]byte `v8:"discManufacturingInfo"`
}

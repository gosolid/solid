//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DVDCopyrightInfo
*/

type DVDCopyrightInfo struct {
  DataLength [2]byte `v8:"dataLength"`
  Reserved [2]byte `v8:"reserved"`
  CopyrightProtectionSystemType byte `v8:"copyrightProtectionSystemType"`
  RegionMask byte `v8:"regionMask"`
  Reserved2 [2]byte `v8:"reserved2"`
}

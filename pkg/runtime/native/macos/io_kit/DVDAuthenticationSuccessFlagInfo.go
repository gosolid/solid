//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DVDAuthenticationSuccessFlagInfo
*/

type DVDAuthenticationSuccessFlagInfo struct {
  DataLength [2]byte `v8:"dataLength"`
  Reserved [2]byte `v8:"reserved"`
  Reserved2 [3]byte `v8:"reserved2"`
  SuccessFlag byte `v8:"successFlag"`
  ReservedBits byte `v8:"reservedBits"`
}

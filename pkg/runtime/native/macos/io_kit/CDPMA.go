//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.CDPMA
*/

type CDPMA struct {
  DataLength uint16 `v8:"dataLength"`
  Reserved byte `v8:"reserved"`
  Reserved2 byte `v8:"reserved2"`
  Descriptors [0]CDPMADescriptor `v8:"descriptors"`
}

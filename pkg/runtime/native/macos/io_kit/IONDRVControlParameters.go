//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IONDRVControlParameters
*/

type IONDRVControlParameters struct {
  ReservedA [26]byte `v8:"reservedA"`
  Code uint16 `v8:"code"`
  Params void `v8:"params"`
  ReservedB [18]byte `v8:"reservedB"`
}

//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IORPCMessageErrorReturnContent
*/

type IORPCMessageErrorReturnContent struct {
  Hdr IORPCMessage `v8:"hdr"`
  Result int `v8:"result"`
  Pad uint `v8:"pad"`
}

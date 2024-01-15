//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOServiceInterestContent64
*/

type IOServiceInterestContent64 struct {
  MessageType uint `v8:"messageType"`
  MessageArgument [1]uint64 `v8:"messageArgument"`
}

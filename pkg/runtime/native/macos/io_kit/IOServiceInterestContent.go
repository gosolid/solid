//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOServiceInterestContent
*/

type IOServiceInterestContent struct {
  MessageType uint `v8:"messageType"`
  MessageArgument [1]*void `v8:"messageArgument"`
}

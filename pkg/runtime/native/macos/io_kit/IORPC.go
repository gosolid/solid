//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IORPC
*/

type IORPC struct {
  Message IORPCMessageMach `v8:"message"`
  Reply IORPCMessageMach `v8:"reply"`
  SendSize uint `v8:"sendSize"`
  ReplySize uint `v8:"replySize"`
}

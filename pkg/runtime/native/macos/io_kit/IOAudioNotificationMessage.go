//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct IOKit.IOAudioNotificationMessage
*/

type IOAudioNotificationMessage struct {
  MessageHeader objc.MachMsgHeaderT `v8:"messageHeader"`
  Type uint `v8:"type"`
  Ref uint `v8:"ref"`
  Sender void `v8:"sender"`
}

//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct IOKit.IODataQueueAppendix
*/

type IODataQueueAppendix struct {
  Version uint `v8:"version"`
  Msgh objc.MachMsgHeaderT `v8:"msgh"`
}

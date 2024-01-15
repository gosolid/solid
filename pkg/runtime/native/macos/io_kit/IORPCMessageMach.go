//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct IOKit.IORPCMessageMach
*/

type IORPCMessageMach struct {
  Msgh objc.MachMsgHeaderT `v8:"msgh"`
  MsghBody objc.MachMsgBodyT `v8:"msghBody"`
  Objects [0]objc.MachMsgPortDescriptorT `v8:"objects"`
}

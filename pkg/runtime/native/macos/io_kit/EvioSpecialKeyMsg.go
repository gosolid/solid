//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct IOKit.evioSpecialKeyMsg
*/

type EvioSpecialKeyMsg struct {
  Head objc.MachMsgHeaderT `v8:"head"`
  Key int `v8:"key"`
  Direction int `v8:"direction"`
  Flags int `v8:"flags"`
  Level int `v8:"level"`
}

//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct IOKit.NXEventExtension
*/

type NXEventExtension struct {
  Flags uint `v8:"flags"`
  Audit objc.AuditTokenT `v8:"audit"`
}

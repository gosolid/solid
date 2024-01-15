//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.OSNotificationHeader64
*/

type OSNotificationHeader64 struct {
  Size uint `v8:"size"`
  Type uint `v8:"type"`
  Reference [8]uint64 `v8:"reference"`
  Content []byte `v8:"content"`
}

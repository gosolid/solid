//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.RegEntryID
*/

type RegEntryID struct {
  Opaque [4]*void `v8:"opaque"`
}

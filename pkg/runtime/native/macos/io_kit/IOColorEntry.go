//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOColorEntry
*/

type IOColorEntry struct {
  Index uint16 `v8:"index"`
  Red uint16 `v8:"red"`
  Green uint16 `v8:"green"`
  Blue uint16 `v8:"blue"`
}

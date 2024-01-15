//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.bm34Cursor
*/

type Bm34Cursor struct {
  Image [4][256]uint16 `v8:"image"`
  Save [256]uint16 `v8:"save"`
}

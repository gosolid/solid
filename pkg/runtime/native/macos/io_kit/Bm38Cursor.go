//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.bm38Cursor
*/

type Bm38Cursor struct {
  Image [4][256]uint `v8:"image"`
  Save [256]uint `v8:"save"`
}

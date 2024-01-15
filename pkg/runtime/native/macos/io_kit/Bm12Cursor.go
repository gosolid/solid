//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.bm12Cursor
*/

type Bm12Cursor struct {
  Image [4][16]uint `v8:"image"`
  Mask [4][16]uint `v8:"mask"`
  Save [16]uint `v8:"save"`
}

//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.bm18Cursor
*/

type Bm18Cursor struct {
  Image [4][256]byte `v8:"image"`
  Mask [4][256]byte `v8:"mask"`
  Save [256]byte `v8:"save"`
}

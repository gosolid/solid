//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOBlitCursor
*/

type IOBlitCursor struct {
  Operation IOBlitOperation `v8:"operation"`
  Rect IOBlitRectangle `v8:"rect"`
}

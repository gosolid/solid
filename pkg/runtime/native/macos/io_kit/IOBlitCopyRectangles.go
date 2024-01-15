//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOBlitCopyRectangles
*/

type IOBlitCopyRectangles struct {
  Operation IOBlitOperation `v8:"operation"`
  Count uint `v8:"count"`
  Rects [1]IOBlitCopyRectangle `v8:"rects"`
}

//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOBlitScanlines
*/

type IOBlitScanlines struct {
  Operation IOBlitOperation `v8:"operation"`
  Count uint `v8:"count"`
  Y int `v8:"y"`
  Height int `v8:"height"`
  X [2]int `v8:"x"`
}

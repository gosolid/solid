//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOAccelDeviceRegion
*/

type IOAccelDeviceRegion struct {
  NumRects uint `v8:"numRects"`
  Bounds IOAccelBounds `v8:"bounds"`
  Rect [0]IOAccelBounds `v8:"rect"`
}

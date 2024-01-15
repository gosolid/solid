//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOBlitCopyRegion
*/

type IOBlitCopyRegion struct {
  Operation IOBlitOperation `v8:"operation"`
  DeltaX int `v8:"deltaX"`
  DeltaY int `v8:"deltaY"`
  Region IOAccelDeviceRegion `v8:"region"`
}

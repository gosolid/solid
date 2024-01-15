//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOAccelSurfaceScaling
*/

type IOAccelSurfaceScaling struct {
  Buffer IOAccelBounds `v8:"buffer"`
  Source IOAccelSize `v8:"source"`
  Reserved [8]uint `v8:"reserved"`
}

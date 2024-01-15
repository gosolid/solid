//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IODisplayModeInformation
*/

type IODisplayModeInformation struct {
  NominalWidth uint `v8:"nominalWidth"`
  NominalHeight uint `v8:"nominalHeight"`
  RefreshRate uint `v8:"refreshRate"`
  MaxDepthIndex int `v8:"maxDepthIndex"`
  Flags uint `v8:"flags"`
  ImageWidth uint16 `v8:"imageWidth"`
  ImageHeight uint16 `v8:"imageHeight"`
  Reserved [3]uint `v8:"reserved"`
}

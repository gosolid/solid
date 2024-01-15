//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOAccelSurfaceInformation
*/

type IOAccelSurfaceInformation struct {
  Address [4]uint64 `v8:"address"`
  RowBytes uint `v8:"rowBytes"`
  Width uint `v8:"width"`
  Height uint `v8:"height"`
  PixelFormat uint `v8:"pixelFormat"`
  Flags uint `v8:"flags"`
  ColorTemperature [4]int `v8:"colorTemperature"`
  TypeDependent [4]uint `v8:"typeDependent"`
}

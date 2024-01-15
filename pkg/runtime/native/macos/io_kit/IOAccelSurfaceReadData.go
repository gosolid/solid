//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOAccelSurfaceReadData
*/

type IOAccelSurfaceReadData struct {
  X int `v8:"x"`
  Y int `v8:"y"`
  W int `v8:"w"`
  H int `v8:"h"`
  ClientAddr uint64 `v8:"clientAddr"`
  ClientRowBytes uint `v8:"clientRowBytes"`
}

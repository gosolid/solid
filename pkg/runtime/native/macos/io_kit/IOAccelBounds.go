//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOAccelBounds
*/

type IOAccelBounds struct {
  X int16 `v8:"x"`
  Y int16 `v8:"y"`
  W int16 `v8:"w"`
  H int16 `v8:"h"`
}

//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOAccelSize
*/

type IOAccelSize struct {
  W int16 `v8:"w"`
  H int16 `v8:"h"`
}

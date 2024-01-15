//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.NXPoint
*/

type NXPoint struct {
  X float32 `v8:"x"`
  Y float32 `v8:"y"`
}

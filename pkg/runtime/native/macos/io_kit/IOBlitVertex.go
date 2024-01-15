//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOBlitVertex
*/

type IOBlitVertex struct {
  X int `v8:"x"`
  Y int `v8:"y"`
}

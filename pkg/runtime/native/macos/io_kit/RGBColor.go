//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.RGBColor
*/

type RGBColor struct {
  Red uint16 `v8:"red"`
  Green uint16 `v8:"green"`
  Blue uint16 `v8:"blue"`
}

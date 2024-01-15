//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.ColorSpec
*/

type ColorSpec struct {
  Value int16 `v8:"value"`
  Rgb RGBColor `v8:"rgb"`
}

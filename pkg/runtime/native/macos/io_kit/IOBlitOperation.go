//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOBlitOperation
*/

type IOBlitOperation struct {
  Color0 uint `v8:"color0"`
  Color1 uint `v8:"color1"`
  OffsetX int `v8:"offsetX"`
  OffsetY int `v8:"offsetY"`
  SourceKeyColor uint `v8:"sourceKeyColor"`
  DestKeyColor uint `v8:"destKeyColor"`
  Specific [16]uint `v8:"specific"`
}

//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOHardwareCursorInfo
*/

type IOHardwareCursorInfo struct {
  MajorVersion uint16 `v8:"majorVersion"`
  MinorVersion uint16 `v8:"minorVersion"`
  CursorHeight uint `v8:"cursorHeight"`
  CursorWidth uint `v8:"cursorWidth"`
  ColorMap IOColorEntry `v8:"colorMap"`
  HardwareCursorData byte `v8:"hardwareCursorData"`
  CursorHotSpotX uint16 `v8:"cursorHotSpotX"`
  CursorHotSpotY uint16 `v8:"cursorHotSpotY"`
  Reserved [5]uint `v8:"reserved"`
}

//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOHardwareCursorDescriptor
*/

type IOHardwareCursorDescriptor struct {
  MajorVersion uint16 `v8:"majorVersion"`
  MinorVersion uint16 `v8:"minorVersion"`
  Height uint `v8:"height"`
  Width uint `v8:"width"`
  BitDepth uint `v8:"bitDepth"`
  MaskBitDepth uint `v8:"maskBitDepth"`
  NumColors uint `v8:"numColors"`
  ColorEncodings uint `v8:"colorEncodings"`
  Flags uint `v8:"flags"`
  SupportedSpecialEncodings uint `v8:"supportedSpecialEncodings"`
  SpecialEncodings [16]uint `v8:"specialEncodings"`
}

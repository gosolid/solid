//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOAudioStreamFormatExtension
*/

type IOAudioStreamFormatExtension struct {
  FVersion uint `v8:"fVersion"`
  FFlags uint `v8:"fFlags"`
  FFramesPerPacket uint `v8:"fFramesPerPacket"`
  FBytesPerPacket uint `v8:"fBytesPerPacket"`
}

//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOAudioStreamDataDescriptor
*/

type IOAudioStreamDataDescriptor struct {
  FVersion uint `v8:"fVersion"`
  FNumberOfStreams uint `v8:"fNumberOfStreams"`
  FStreamLength [1]uint `v8:"fStreamLength"`
}

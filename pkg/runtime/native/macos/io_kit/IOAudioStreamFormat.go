//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOAudioStreamFormat
*/

type IOAudioStreamFormat struct {
  FNumChannels uint `v8:"fNumChannels"`
  FSampleFormat uint `v8:"fSampleFormat"`
  FNumericRepresentation uint `v8:"fNumericRepresentation"`
  FBitDepth byte `v8:"fBitDepth"`
  FBitWidth byte `v8:"fBitWidth"`
  FAlignment byte `v8:"fAlignment"`
  FByteOrder byte `v8:"fByteOrder"`
  FIsMixable byte `v8:"fIsMixable"`
  FDriverTag uint `v8:"fDriverTag"`
}

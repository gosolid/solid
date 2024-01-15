//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.CDTEXTDescriptor
*/

type CDTEXTDescriptor struct {
  PackType byte `v8:"packType"`
  TrackNumber byte `v8:"trackNumber"`
  SequenceNumber byte `v8:"sequenceNumber"`
  CharacterPosition byte `v8:"characterPosition"`
  BlockNumber byte `v8:"blockNumber"`
  DoubleByteCharacterCode byte `v8:"doubleByteCharacterCode"`
  TextData [12]byte `v8:"textData"`
  Reserved [2]byte `v8:"reserved"`
}

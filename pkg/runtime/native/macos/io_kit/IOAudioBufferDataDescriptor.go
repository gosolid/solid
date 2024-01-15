//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOAudioBufferDataDescriptor
*/

type IOAudioBufferDataDescriptor struct {
  FActualDataByteSize uint `v8:"fActualDataByteSize"`
  FActualNumSampleFrames uint `v8:"fActualNumSampleFrames"`
  FTotalDataByteSize uint `v8:"fTotalDataByteSize"`
  FNominalDataByteSize uint `v8:"fNominalDataByteSize"`
  FData [1]byte `v8:"fData"`
}

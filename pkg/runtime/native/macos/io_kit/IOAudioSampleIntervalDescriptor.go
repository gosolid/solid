//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOAudioSampleIntervalDescriptor
*/

type IOAudioSampleIntervalDescriptor struct {
  SampleIntervalHi uint `v8:"sampleIntervalHi"`
  SampleIntervalLo uint `v8:"sampleIntervalLo"`
}

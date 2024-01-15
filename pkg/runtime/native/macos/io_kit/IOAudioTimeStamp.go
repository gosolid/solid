//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOAudioTimeStamp
*/

type IOAudioTimeStamp struct {
  FSampleTime uint64 `v8:"fSampleTime"`
  FHostTime uint64 `v8:"fHostTime"`
  FRateScalar uint64 `v8:"fRateScalar"`
  FWordClockTime uint64 `v8:"fWordClockTime"`
  FSMPTETime IOAudioSMPTETime `v8:"fSMPTETime"`
  FFlags uint `v8:"fFlags"`
  FReserved uint `v8:"fReserved"`
}

//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOAudioSMPTETime
*/

type IOAudioSMPTETime struct {
  FSubframes int16 `v8:"fSubframes"`
  FSubframeDivisor int16 `v8:"fSubframeDivisor"`
  FCounter uint `v8:"fCounter"`
  FType uint `v8:"fType"`
  FFlags uint `v8:"fFlags"`
  FHours int16 `v8:"fHours"`
  FMinutes int16 `v8:"fMinutes"`
  FSeconds int16 `v8:"fSeconds"`
  FFrames int16 `v8:"fFrames"`
}

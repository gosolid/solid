//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.CommonChunk
*/

type CommonChunk struct {
  CkID uint `v8:"ckID"`
  CkSize int `v8:"ckSize"`
  NumChannels int16 `v8:"numChannels"`
  NumSampleFrames uint `v8:"numSampleFrames"`
  SampleSize int16 `v8:"sampleSize"`
  SampleRate objc.Float80 `v8:"sampleRate"`
}

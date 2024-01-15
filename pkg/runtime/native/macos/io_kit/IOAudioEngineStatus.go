//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct IOKit.IOAudioEngineStatus
*/

type IOAudioEngineStatus struct {
  FVersion uint `v8:"fVersion"`
  FCurrentLoopCount uint `v8:"fCurrentLoopCount"`
  FLastLoopTime objc.UnsignedWide `v8:"fLastLoopTime"`
  FEraseHeadSampleFrame uint `v8:"fEraseHeadSampleFrame"`
}

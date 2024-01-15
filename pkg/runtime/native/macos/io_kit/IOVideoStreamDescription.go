//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOVideoStreamDescription
*/

type IOVideoStreamDescription struct {
  MVideoCodecType uint `v8:"mVideoCodecType"`
  MVideoCodecFlags uint `v8:"mVideoCodecFlags"`
  MWidth uint `v8:"mWidth"`
  MHeight uint `v8:"mHeight"`
  MReserved1 uint `v8:"mReserved1"`
  MReserved2 uint `v8:"mReserved2"`
}

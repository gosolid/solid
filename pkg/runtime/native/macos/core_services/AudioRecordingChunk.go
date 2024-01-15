//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.AudioRecordingChunk
*/

type AudioRecordingChunk struct {
  CkID uint `v8:"ckID"`
  CkSize int `v8:"ckSize"`
  AESChannelStatus [24]byte `v8:"aESChannelStatus"`
}

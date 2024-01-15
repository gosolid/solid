//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.SoundDataChunk
*/

type SoundDataChunk struct {
  CkID uint `v8:"ckID"`
  CkSize int `v8:"ckSize"`
  Offset uint `v8:"offset"`
  BlockSize uint `v8:"blockSize"`
}

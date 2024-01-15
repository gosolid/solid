//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.MIDIDataChunk
*/

type MIDIDataChunk struct {
  CkID uint `v8:"ckID"`
  CkSize int `v8:"ckSize"`
  MIDIdata [1]byte `v8:"mIDIdata"`
}

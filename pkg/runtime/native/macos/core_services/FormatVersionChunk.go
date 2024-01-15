//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FormatVersionChunk
*/

type FormatVersionChunk struct {
  CkID uint `v8:"ckID"`
  CkSize int `v8:"ckSize"`
  Timestamp uint `v8:"timestamp"`
}

//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ApplicationSpecificChunk
*/

type ApplicationSpecificChunk struct {
  CkID uint `v8:"ckID"`
  CkSize int `v8:"ckSize"`
  ApplicationSignature uint `v8:"applicationSignature"`
  Data [1]byte `v8:"data"`
}

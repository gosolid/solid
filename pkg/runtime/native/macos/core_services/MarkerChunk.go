//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.MarkerChunk
*/

type MarkerChunk struct {
  CkID uint `v8:"ckID"`
  CkSize int `v8:"ckSize"`
  NumMarkers uint16 `v8:"numMarkers"`
  Markers [1]Marker `v8:"markers"`
}

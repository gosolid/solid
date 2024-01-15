//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.Marker
*/

type Marker struct {
  Id int16 `v8:"id"`
  Position uint `v8:"position"`
  MarkerName [256]byte `v8:"markerName"`
}

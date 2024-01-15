//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.Comment
*/

type Comment struct {
  TimeStamp uint `v8:"timeStamp"`
  Marker int16 `v8:"marker"`
  Count uint16 `v8:"count"`
  Text [1]byte `v8:"text"`
}

//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECSubTextEncodingsRec
*/

type TECSubTextEncodingsRec struct {
  Count uint `v8:"count"`
  SubTextEncodingRec TECSubTextEncodingRec `v8:"subTextEncodingRec"`
}

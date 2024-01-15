//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.UnicodeMapping
*/

type UnicodeMapping struct {
  UnicodeEncoding uint `v8:"unicodeEncoding"`
  OtherEncoding uint `v8:"otherEncoding"`
  MappingVersion int `v8:"mappingVersion"`
}

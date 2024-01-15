//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECLocaleListToEncodingListRec
*/

type TECLocaleListToEncodingListRec struct {
  Offset uint `v8:"offset"`
  Count uint `v8:"count"`
  Locales int16 `v8:"locales"`
}

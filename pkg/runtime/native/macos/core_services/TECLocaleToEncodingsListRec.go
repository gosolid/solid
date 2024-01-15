//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECLocaleToEncodingsListRec
*/

type TECLocaleToEncodingsListRec struct {
  Count uint `v8:"count"`
  LocaleListToEncodingList TECLocaleListToEncodingListRec `v8:"localeListToEncodingList"`
}

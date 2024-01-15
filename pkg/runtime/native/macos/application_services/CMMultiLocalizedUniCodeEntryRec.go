//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMMultiLocalizedUniCodeEntryRec
*/

type CMMultiLocalizedUniCodeEntryRec struct {
  LanguageCode [2]byte `v8:"languageCode"`
  RegionCode [2]byte `v8:"regionCode"`
  TextLength uint `v8:"textLength"`
  TextOffset uint `v8:"textOffset"`
}

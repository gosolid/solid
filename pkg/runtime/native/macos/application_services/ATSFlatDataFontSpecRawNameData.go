//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFlatDataFontSpecRawNameData
*/

type ATSFlatDataFontSpecRawNameData struct {
  FontNameType uint `v8:"fontNameType"`
  FontNamePlatform uint `v8:"fontNamePlatform"`
  FontNameScript uint `v8:"fontNameScript"`
  FontNameLanguage uint `v8:"fontNameLanguage"`
  FontNameLength uint `v8:"fontNameLength"`
}

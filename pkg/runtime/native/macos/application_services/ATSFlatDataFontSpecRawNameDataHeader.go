//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFlatDataFontSpecRawNameDataHeader
*/

type ATSFlatDataFontSpecRawNameDataHeader struct {
  NumberOfFlattenedNames uint `v8:"numberOfFlattenedNames"`
  NameDataArray [1]ATSFlatDataFontSpecRawNameData `v8:"nameDataArray"`
}

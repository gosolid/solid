//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFlatDataFontNameDataHeader
*/

type ATSFlatDataFontNameDataHeader struct {
  NameSpecifierType uint `v8:"nameSpecifierType"`
  NameSpecifierSize uint `v8:"nameSpecifierSize"`
}

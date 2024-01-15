//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFlatDataLineInfoHeader
*/

type ATSFlatDataLineInfoHeader struct {
  NumberOfLines uint `v8:"numberOfLines"`
  LineInfoArray [1]ATSFlatDataLineInfoData `v8:"lineInfoArray"`
}

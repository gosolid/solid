//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFlatDataStyleListStyleDataHeader
*/

type ATSFlatDataStyleListStyleDataHeader struct {
  SizeOfStyleInfo uint `v8:"sizeOfStyleInfo"`
  NumberOfSetAttributes uint `v8:"numberOfSetAttributes"`
  NumberOfSetFeatures uint `v8:"numberOfSetFeatures"`
  NumberOfSetVariations uint `v8:"numberOfSetVariations"`
}

//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFlatDataStyleListHeader
*/

type ATSFlatDataStyleListHeader struct {
  NumberOfStyles uint `v8:"numberOfStyles"`
  StyleDataArray [1]ATSFlatDataStyleListStyleDataHeader `v8:"styleDataArray"`
}

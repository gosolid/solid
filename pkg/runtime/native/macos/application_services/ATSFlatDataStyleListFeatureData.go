//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFlatDataStyleListFeatureData
*/

type ATSFlatDataStyleListFeatureData struct {
  TheFeatureType uint16 `v8:"theFeatureType"`
  TheFeatureSelector uint16 `v8:"theFeatureSelector"`
}

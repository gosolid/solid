//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFlatDataStyleListVariationData
*/

type ATSFlatDataStyleListVariationData struct {
  TheVariationAxis uint `v8:"theVariationAxis"`
  TheVariationValue int `v8:"theVariationValue"`
}

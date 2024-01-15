//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFlatDataLineInfoData
*/

type ATSFlatDataLineInfoData struct {
  LineLength uint `v8:"lineLength"`
  NumberOfLineControls uint `v8:"numberOfLineControls"`
}

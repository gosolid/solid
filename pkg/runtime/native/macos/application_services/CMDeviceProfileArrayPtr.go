//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMDeviceProfileArrayPtr
*/

type CMDeviceProfileArrayPtr struct {
  ProfileCount uint `v8:"profileCount"`
  Profiles [1]CMDeviceProfileInfo `v8:"profiles"`
}

//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSURGBAlphaColor
*/

type ATSURGBAlphaColor struct {
  Red float32 `v8:"red"`
  Green float32 `v8:"green"`
  Blue float32 `v8:"blue"`
  Alpha float32 `v8:"alpha"`
}

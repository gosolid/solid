//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMLuvColor
*/

type CMLuvColor struct {
  L uint16 `v8:"l"`
  U uint16 `v8:"u"`
  V uint16 `v8:"v"`
}

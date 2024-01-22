//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMRGBColor
*/

type CMRGBColor struct {
  Red uint16 `v8:"red"`
  Green uint16 `v8:"green"`
  Blue uint16 `v8:"blue"`
}
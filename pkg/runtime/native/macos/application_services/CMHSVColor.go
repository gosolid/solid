//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMHSVColor
*/

type CMHSVColor struct {
  Hue uint16 `v8:"hue"`
  Saturation uint16 `v8:"saturation"`
  Value uint16 `v8:"value"`
}

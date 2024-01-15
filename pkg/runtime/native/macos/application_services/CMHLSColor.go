//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMHLSColor
*/

type CMHLSColor struct {
  Hue uint16 `v8:"hue"`
  Lightness uint16 `v8:"lightness"`
  Saturation uint16 `v8:"saturation"`
}

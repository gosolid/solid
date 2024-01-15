//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMYxyColor
*/

type CMYxyColor struct {
  CapY uint16 `v8:"capY"`
  X uint16 `v8:"x"`
  Y uint16 `v8:"y"`
}

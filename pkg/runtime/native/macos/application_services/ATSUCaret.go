//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSUCaret
*/

type ATSUCaret struct {
  FX int `v8:"fX"`
  FY int `v8:"fY"`
  FDeltaX int `v8:"fDeltaX"`
  FDeltaY int `v8:"fDeltaY"`
}

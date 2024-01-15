//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMXYZColor
*/

type CMXYZColor struct {
  X uint16 `v8:"x"`
  Y uint16 `v8:"y"`
  Z uint16 `v8:"z"`
}

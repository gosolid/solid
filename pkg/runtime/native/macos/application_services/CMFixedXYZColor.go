//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMFixedXYZColor
*/

type CMFixedXYZColor struct {
  X int `v8:"x"`
  Y int `v8:"y"`
  Z int `v8:"z"`
}

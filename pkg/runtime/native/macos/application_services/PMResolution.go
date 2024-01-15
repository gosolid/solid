//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.PMResolution
*/

type PMResolution struct {
  HRes float64 `v8:"hRes"`
  VRes float64 `v8:"vRes"`
}

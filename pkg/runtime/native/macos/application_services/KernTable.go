//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.KernTable
*/

type KernTable struct {
  NumKerns int16 `v8:"numKerns"`
}

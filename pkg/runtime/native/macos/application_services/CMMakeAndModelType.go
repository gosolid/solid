//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMMakeAndModelType
*/

type CMMakeAndModelType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  MakeAndModel CMMakeAndModel `v8:"makeAndModel"`
}

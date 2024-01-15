//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMS15Fixed16ArrayType
*/

type CMS15Fixed16ArrayType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  Value [1]int `v8:"value"`
}

//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMU16Fixed16ArrayType
*/

type CMU16Fixed16ArrayType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  Value [1]uint `v8:"value"`
}

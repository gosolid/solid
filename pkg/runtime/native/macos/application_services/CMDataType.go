//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMDataType
*/

type CMDataType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  DataFlag uint `v8:"dataFlag"`
  Data [1]byte `v8:"data"`
}

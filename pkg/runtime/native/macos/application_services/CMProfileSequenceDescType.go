//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMProfileSequenceDescType
*/

type CMProfileSequenceDescType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  Count uint `v8:"count"`
  Data [1]byte `v8:"data"`
}

//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMTextDescriptionType
*/

type CMTextDescriptionType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  ASCIICount uint `v8:"aSCIICount"`
  ASCIIName [2]byte `v8:"aSCIIName"`
}

//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMUnicodeTextType
*/

type CMUnicodeTextType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  Text [1]uint16 `v8:"text"`
}

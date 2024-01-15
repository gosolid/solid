//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMUcrBgType
*/

type CMUcrBgType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  UcrCount uint `v8:"ucrCount"`
  UcrValues [1]uint16 `v8:"ucrValues"`
}

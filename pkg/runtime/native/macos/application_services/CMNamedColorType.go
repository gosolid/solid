//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMNamedColorType
*/

type CMNamedColorType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  VendorFlag uint `v8:"vendorFlag"`
  Count uint `v8:"count"`
  PrefixName [1]byte `v8:"prefixName"`
}
